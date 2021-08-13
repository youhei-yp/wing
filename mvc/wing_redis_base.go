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
	"github.com/garyburd/redigo/redis"
	"github.com/youhei-yp/wing/invar"
	"github.com/youhei-yp/wing/logger"
)

// Set set a value of a key.
//
//	option[0] : OptEX, OptPX, OptEXAT, OptPXAT, OptNX, OptXX, OptKEEPTTL
//	option[1] : time number in seconds, milliseconds, or unix timestampe
//
//	[CODE:]
//	err := c.Set("foo1", 123456)
//	err := c.Set("foo2", 123456, OptEX, 60 * 10)
//	err := c.Set("foo3", 56.789, OptPX, 1000000)
//	err := c.Set("foo4", "5678", OptEXAT, 1628778035)
//	err := c.Set("foo5", true,   OptPXAT, 1628778035)
//	err := c.Set("foo6", 111333, OptNX)
//	err := c.Set("foo7", 222444, OptXX)
//	err := c.Set("foo8", 333555, OptKEEPTTL)
//	[CODE]
//
//	see https://redis.io/commands/set
func (c *WingRedisConn) Set(key string, value interface{}, options ...interface{}) error {
	err := invar.ErrInvalidRedisOptions
	if options != nil && len(options) > 0 {
		switch option := options[0].(type) {
		case string:
			if len(options) > 1 {
				switch expire := options[1].(type) {
				case int64:
					switch option {
					case OptEX, OptPX, OptEXAT, OptPXAT:
						_, err = c.Conn.Do("SET", c.serviceNamespace+key, value, option, expire)
					}
				}
			} else {
				switch option {
				case OptNX, OptXX, OptKEEPTTL:
					_, err = c.Conn.Do("SET", c.serviceNamespace+key, value, option)
				}
			}
		}
	} else {
		_, err = c.Conn.Do("SET", c.serviceNamespace+key, value)
	}
	return err
}

// SetEx set a value and expiration in seconds of a key.
func (c *WingRedisConn) SetEx(key string, value interface{}, expire int64) error {
	return c.setWithExpire(key, "SETEX", value, expire)
}

// SetPx set a value and expiration in milliseconds of a key.
func (c *WingRedisConn) SetPx(key string, value interface{}, expire int64) error {
	return c.setWithExpire(key, "PSETEX", value, expire)
}

// SetNx set a value of a unexist key, it return false when failed set on exist key.
//
//	see https://redis.io/commands/setnx
func (c *WingRedisConn) SetNx(key string, value interface{}) (bool, error) {
	exist, err := c.Conn.Do("SETNX", c.serviceNamespace+key, value)
	if err != nil {
		return false, err
	}
	return (exist == 1), nil
}

// SetRange overwrite part of a string at key starting at the specified offset,
// it will transform the given value to string first, and retuen the memery length.
//
//	see https://redis.io/commands/setrange
func (c *WingRedisConn) SetRange(key string, value interface{}, offset int) int {
	length, err := redis.Int(c.Conn.Do("SETRANGE", c.serviceNamespace+key, offset, value))
	if err != nil {
		logger.E("Redis:SETRANGE [key"+key+"] err:", err)
		return -1
	}
	return length
}

// Append append a value of a key, it will transform the given value to string first,
// than append end of exist string, or set value same as SET commond.
//
//	see https://redis.io/commands/append
func (c *WingRedisConn) Append(key string, value interface{}) int {
	length, err := redis.Int(c.Conn.Do("APPEND", c.serviceNamespace+key, value))
	if err != nil {
		logger.E("Redis:APPEND [key"+key+"] err:", err)
		return -1
	}
	return length
}

// StrLength get the string value of key, it return 0 if the key unexist.
//
//	see https://redis.io/commands/strlen
func (c *WingRedisConn) StrLength(key string) int {
	length, _ := redis.Int(c.Conn.Do("STRLEN", c.serviceNamespace+key))
	return length
}

// Exist determine if a key exists, only support one key one check, set Exists.
func (c *WingRedisConn) Exist(key string) (bool, error) {
	return redis.Bool(c.Conn.Do("EXISTS", c.serviceNamespace+key))
}

// Exists determine if given keys exist, and return exist keys count,
// you can use NsKey() or Nskeys() to transform origin keys to namespaced keys,
// and then call Exists(nskeys[0], nskeys[1] ... nskeys[m]) to multiple check.
//
//	see https://redis.io/commands/exists
func (c *WingRedisConn) Exists(nskeys ...string) (int, error) {
	return redis.Int(c.Conn.Do("EXISTS", nskeys))
}

// Expire set a key's time to live in seconds, the optional values can be set
// as ExpNX, ExpXX, ExpGT, ExpLT since Redis 7.0 support.
func (c *WingRedisConn) Expire(key string, expire int64, option ...string) bool {
	return c.setKeyExpire(key, "EXPIRE", expire, option...)
}

// ExpireAt set the expiration for a key as a unix timestamp in seconds, the optional
// values can be set as ExpNX, ExpXX, ExpGT, ExpLT since Redis 7.0 support.
func (c *WingRedisConn) ExpireAt(key string, expire int64, option ...string) bool {
	return c.setKeyExpire(key, "EXPIREAT", expire, option...)
}

// Persist remove the expiration from a key.
//
//	see https://redis.io/commands/persist
func (c *WingRedisConn) Persist(key string) bool {
	set, err := redis.Bool(c.Conn.Do("PERSIST", c.serviceNamespace+key))
	if err != nil {
		logger.E("Redis:PERSIST [key"+key+"] err:", err)
		return false
	}
	return set
}

// GetExpire get the time to live for a key in secons.
func (c *WingRedisConn) GetExpire(key string) (int64, error) {
	return c.getKeyExpire(key, "TTL")
}

// GetExpireMs get the time to live for a key in milliseconds.
func (c *WingRedisConn) GetExpireMs(key string) (int64, error) {
	return c.getKeyExpire(key, "PTTL")
}

// Delete delete a key, see Deletes.
func (c *WingRedisConn) Delete(key string) bool {
	deleted, err := redis.Bool(c.Conn.Do("DEL", c.serviceNamespace+key))
	if err != nil {
		logger.E("Redis:DEL [key"+key+"] err:", err)
		return false
	}
	return deleted
}

// Deletes delete the given keys, and return deleted keys count.
// you can use NsKey() or Nskeys() to transform origin keys to namespaced keys,
// and then call Deletes(nskeys[0], nskeys[1] ... nskeys[m]) to multiple delete.
//
//	see https://redis.io/commands/del
func (c *WingRedisConn) Deletes(nskeys ...string) (int, error) {
	return redis.Int(c.Conn.Do("DEL", nskeys))
}

// GetRange get the string value of key cut by given range.
//
//	see https://redis.io/commands/getrange
func (c *WingRedisConn) GetRange(key string, start, end int) (string, error) {
	return redis.String(c.Conn.Do("GETRANGE", c.serviceNamespace+key, start, end))
}

// Get get the string value of key.
func (c *WingRedisConn) Get(key string) (string, error) {
	return redis.String(c.getWithOptions(key))
}

// GetInt get the int value of key.
func (c *WingRedisConn) GetInt(key string) (int, error) {
	return redis.Int(c.getWithOptions(key))
}

// GetInt64 get the int64 value of key.
func (c *WingRedisConn) GetInt64(key string) (int64, error) {
	return redis.Int64(c.getWithOptions(key))
}

// GetUint64 get the uint64 value of key.
func (c *WingRedisConn) GetUint64(key string) (uint64, error) {
	return redis.Uint64(c.getWithOptions(key))
}

// GetFloat64 get the float value of key.
func (c *WingRedisConn) GetFloat64(key string) (float64, error) {
	return redis.Float64(c.getWithOptions(key))
}

// GetBytes get the bytes array of key.
func (c *WingRedisConn) GetBytes(key string) ([]byte, error) {
	return redis.Bytes(c.getWithOptions(key))
}

// GetBool get the bool value of key.
func (c *WingRedisConn) GetBool(key string) (bool, error) {
	return redis.Bool(c.getWithOptions(key))
}

// GetDel get and delete the string value of key
func (c *WingRedisConn) GetDel(key string) (string, error) {
	return redis.String(c.getWithOptions(key, CusOptDel))
}

// GetDelInt get and delete the int value of key.
func (c *WingRedisConn) GetDelInt(key string) (int, error) {
	return redis.Int(c.getWithOptions(key, CusOptDel))
}

// GetDelInt64 get and delete the int64 value of key.
func (c *WingRedisConn) GetDelInt64(key string) (int64, error) {
	return redis.Int64(c.getWithOptions(key, CusOptDel))
}

// GetDelUint64 get and delete the uint64 value of key.
func (c *WingRedisConn) GetDelUint64(key string) (uint64, error) {
	return redis.Uint64(c.getWithOptions(key, CusOptDel))
}

// GetDelFloat64 get and delete the float value of key, see Pull.
func (c *WingRedisConn) GetDelFloat64(key string) (float64, error) {
	return redis.Float64(c.getWithOptions(key, CusOptDel))
}

// GetDelBytes get and delete the bytes array of key, see Pull.
func (c *WingRedisConn) GetDelBytes(key string) ([]byte, error) {
	return redis.Bytes(c.getWithOptions(key, CusOptDel))
}

// GetDelBool get and delete the bool value of key, see Pull.
func (c *WingRedisConn) GetDelBool(key string) (bool, error) {
	return redis.Bool(c.getWithOptions(key, CusOptDel))
}

// GETEX get the string value of a key and optionally set its expiration.
func (c *WingRedisConn) GetEx(key string, option string, expire int64) (string, error) {
	return redis.String(c.getWithOptions(key, option, expire))
}

// GetExInt get the int value of key and optionally set its expiration.
func (c *WingRedisConn) GetExInt(key string, option string, expire int64) (int, error) {
	return redis.Int(c.getWithOptions(key, option, expire))
}

// GetExInt64 get the int64 value of key and optionally set its expiration.
func (c *WingRedisConn) GetExInt64(key string, option string, expire int64) (int64, error) {
	return redis.Int64(c.getWithOptions(key, option, expire))
}

// GetExUint64 get the uint64 value of key and optionally set its expiration.
func (c *WingRedisConn) GetExUint64(key string, option string, expire int64) (uint64, error) {
	return redis.Uint64(c.getWithOptions(key, option, expire))
}

// GetExFloat64 get the float value of key and optionally set its expiration.
func (c *WingRedisConn) GetExFloat64(key string, option string, expire int64) (float64, error) {
	return redis.Float64(c.getWithOptions(key, option, expire))
}

// GetExBytes get the bytes array of key and optionally set its expiration.
func (c *WingRedisConn) GetExBytes(key string, option string, expire int64) ([]byte, error) {
	return redis.Bytes(c.getWithOptions(key, option, expire))
}

// GetExBool get the bool value of key and optionally set its expiration.
func (c *WingRedisConn) GetExBool(key string, option string, expire int64) (bool, error) {
	return redis.Bool(c.getWithOptions(key, option, expire))
}
