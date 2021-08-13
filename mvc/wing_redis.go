// Copyright (c) 2018-2022 WING All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2021/08/11   yangping       New version
// -------------------------------------------------------------------

package mvc

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"github.com/youhei-yp/wing/invar"
	"github.com/youhei-yp/wing/logger"
)

// WingRedisConn content provider to support redis utils
type WingRedisConn struct {
	Conn redis.Conn

	// serviceNamespace service namespace to distinguish others server, it will
	// append with key as prefix to get or set data to redis database, default is empty.
	//
	// NOTICE :
	// you must config namespace in /conf/app.config file, or call WingRedis.SetNamespace(ns)
	// to set unique server namespace when multiple services connecting the one redis server.
	serviceNamespace string

	// deadlockDur deadlock max duration
	deadlockDuration int64
}

const (
	redisConfigHost = "redis::host"      // configs key of redis host and port
	redisConfigPwd  = "redis::pwd"       // configs key of redis password
	redisConfigNs   = "redis::namespace" // configs key of redis namespace
	redisConfigLock = "redis::deadlock"  // configs key of redis lock max duration
)

// The follow options may support by diffrent Redis version, get more infor
// by link https://redis.io/commands webset.
const (
	OptEX      = "EX"      // seconds -- Set the specified expire time, in seconds
	OptPX      = "PX"      // milliseconds -- Set the specified expire time, in milliseconds.
	OptEXAT    = "EXAT"    // timestamp-seconds -- Set the specified Unix time at which the key will expire, in seconds.
	OptPXAT    = "PXAT"    // timestamp-milliseconds -- Set the specified Unix time at which the key will expire, in milliseconds.
	OptNX      = "NX"      // Only set the key if it does not already exist.
	OptXX      = "XX"      // Only set the key if it already exist.
	OptKEEPTTL = "KEEPTTL" // Retain the time to live associated with the key, use for SET commond.
	ExpNX      = "NX"      // Set expiry only when the key has no expiry.
	ExpXX      = "XX"      // Set expiry only when the key has an existing expiry.
	ExpGT      = "GT"      // Set expiry only when the new expiry is greater than current one.
	ExpLT      = "LT"      // Set expiry only when the new expiry is less than current one.

	CusOptDel = "DELETE" // The custom option to delete redis key after get commond execute.
)

var (
	// WingRedis a connecter to access redis database data,
	// it will nil before mvc.OpenRedis() called
	WingRedis *WingRedisConn
)

// readRedisCofnigs read redis params from config file, than verify them if empty.
func readRedisCofnigs() (string, string, string, int64, error) {
	host := beego.AppConfig.String(redisConfigHost)
	pwd := beego.AppConfig.String(redisConfigPwd)
	ns := beego.AppConfig.String(redisConfigNs) // allow empty
	if host == "" || pwd == "" {
		return "", "", "", 0, invar.ErrInvalidConfigs
	}

	lock, _ := beego.AppConfig.Int64(redisConfigLock)
	if lock <= 0 {
		lock = 20 // default 20 seconds
	}
	return host, pwd, ns, lock, nil
}

// OpenRedis connect redis database server and auth password,
// the connections holded by mvc.WingRedis object.
//
// NOTICE : you must config redis params in /conf/app.config file as:
//	~
//	[redis]
//	host = "127.0.0.1:6379"
//	pwd = "123456"
//	namespace = "project_namespace"
//	deadlock = 20
//	~
//
//	host is the redis server host ip and port.
//	pwd is the redis server authenticate password.
//	namespace is the prefix string or store key.
//	deadlock is the max time of deadlock, in seconds.
func OpenRedis() error {
	host, pwd, ns, lock, err := readRedisCofnigs()
	if err != nil {
		return err
	}

	// dial TCP connection
	con, err := redis.Dial("tcp", host)
	if err != nil {
		return err
	}

	// authenticate connection password.
	// see https://redis.io/commands/auth
	if _, err = con.Do("AUTH", pwd); err != nil {
		return err
	}

	WingRedis = &WingRedisConn{con, ns, lock}
	return nil
}

// SetNamespace set server uniqu namespace
func (c *WingRedisConn) SetNamespace(ns string) {
	if ns != "" {
		c.serviceNamespace = ns
	}
}

// SetDeadlock set the max deadlock duration
func (c *WingRedisConn) SetDeadlock(dur int64) {
	if dur > 0 {
		c.deadlockDuration = dur
	}
}

// GetDeadlock get the max deadlock duration
func (c *WingRedisConn) GetDeadlock() int64 {
	return c.deadlockDuration
}

// NsKey transform origin key to namespaced key
func (c *WingRedisConn) NsKey(key string) string {
	return c.serviceNamespace + key
}

// NsKeys transform origin keys to namespaced keys
func (c *WingRedisConn) NsKeys(keys ...string) []string {
	return c.NsArrKeys(keys)
}

// NsArrKeys transform origin keys to namespaced keys
func (c *WingRedisConn) NsArrKeys(keys []string) []string {
	nskeys := []string{}
	if keys != nil && len(keys) > 0 {
		for _, key := range keys {
			nskeys = append(nskeys, key)
		}
	}
	return nskeys
}

// ServerTime get redis server unix time is seconds and microsecodns.
//
//	see https://redis.io/commands/time
func (c *WingRedisConn) ServerTime() (int64, int64) {
	st, err := redis.Int64s(c.Conn.Do("TIME"))
	if err != nil || len(st) != 2 {
		logger.E("Redis:TIME Failed set redis server time")
		return 0, 0
	}
	return st[0], st[1]
}

// setWithExpire set a value and expiration of a key.
//
//	see https://redis.io/commands/setex
//	see https://redis.io/commands/psetex
func (c *WingRedisConn) setWithExpire(key, commond string, value interface{}, expire int64) error {
	if _, err := c.Conn.Do(commond, c.serviceNamespace+key, expire, value); err != nil {
		return err
	}
	return nil
}

// parseGetOptions parse the GETEX, GETDEL commonds options
func (c *WingRedisConn) parseGetOptions(options ...interface{}) (string, int64) {
	if options == nil || len(options) == 0 {
		logger.E("Redis: invalid options, parse failed")
		return "", 0
	}

	switch option := options[0].(type) {
	case string:
		if len(options) > 1 /* parse expire */ {
			switch expire := options[1].(type) {
			case int64:
				return option, expire
			default:
				logger.E("Redis: expire value type is not int64")
			}
		} else {
			return option, 0 // just for CusOptDel
		}
	default:
		logger.E("Redis: option value type is not string")
	}
	return "", 0
}

// Get get a value of key, than set value expire time or delete by given options.
//
//	see https://redis.io/commands/get
//	see https://redis.io/commands/getex
//	see https://redis.io/commands/getdel
func (c *WingRedisConn) getWithOptions(key string, options ...interface{}) (interface{}, error) {
	if options != nil {
		var reply interface{}
		err := invar.ErrInvalidRedisOptions
		if option, expire := c.parseGetOptions(options); option != "" {
			switch option {
			case CusOptDel:
				reply, err = c.Conn.Do("GETDEL", c.serviceNamespace+key)
			case OptEX, OptPX, OptEXAT, OptPXAT:
				reply, err = c.Conn.Do("GETEX", c.serviceNamespace+key, option, expire)
			}
		}
		return redis.String(reply, err)
	}
	return c.Conn.Do("GET", c.serviceNamespace+key)
}

// getKeyExpire get the time to live for a key by given commond, it may return unexist
// key error if the key unexist or set expiration and now expire, or return keeplive
// error if the exist key has no associated expire.
//
//	see https://redis.io/commands/ttl
//	see https://redis.io/commands/pttl
func (c *WingRedisConn) getKeyExpire(key, commond string) (int64, error) {
	expire, err := redis.Int64(c.Conn.Do(commond, c.serviceNamespace+key))
	if expire == -2 {
		return 0, invar.ErrUnexistRedisKey
	} else if expire == -1 {
		return 0, invar.ErrNoAssociatedExpire
	}
	return expire, err
}

// setKeyExpire set the expiration for a key by given commond, the optional values
// can be set as ExpNX, ExpXX, ExpGT, ExpLT since Redis 7.0 support.
//
//	see https://redis.io/commands/expire
//	see https://redis.io/commands/expireat
func (c *WingRedisConn) setKeyExpire(key, commond string, expire int64, option ...string) bool {
	set, err := false, invar.ErrInvalidRedisOptions
	if option != nil && len(option) > 0 {
		switch option[0] {
		case ExpNX, ExpXX, ExpGT, ExpLT:
			set, err = redis.Bool(c.Conn.Do(commond, c.serviceNamespace+key, expire, option[0]))
		}
	} else {
		set, err = redis.Bool(c.Conn.Do(commond, c.serviceNamespace+key, expire))
	}

	if err != nil {
		logger.E("Redis:EXPIRE [key"+key+":"+commond+"] err:", err)
		return false
	}
	return set
}
