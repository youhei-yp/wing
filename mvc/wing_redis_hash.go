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
	"github.com/youhei-yp/wing/logger"
)

// HSet set the string value of a hash field.
func (c *WingRedisConn) HSet(key string, field, value interface{}) bool {
	set, err := redis.Bool(c.Conn.Do("HSET", c.serviceNamespace+key, field, value))
	if err != nil {
		logger.E("Redis:HSET [key"+key+"] err:", err)
		return false
	}
	return set
}

// HGet get the string value of a hash field.
func (c *WingRedisConn) HGet(key string, field interface{}) (string, error) {
	return redis.String(c.Conn.Do("HGET", c.serviceNamespace+key, field))
}

// HGetInt get the int value of a hash field.
func (c *WingRedisConn) HGetInt(key string, field interface{}) (int, error) {
	return redis.Int(c.Conn.Do("HGET", c.serviceNamespace+key, field))
}

// HGetInt64 get the int64 value of a hash field.
func (c *WingRedisConn) HGetInt64(key string, field interface{}) (int64, error) {
	return redis.Int64(c.Conn.Do("HGET", c.serviceNamespace+key, field))
}

// HGetUint64 get the uint64 value of a hash field.
func (c *WingRedisConn) HGetUint64(key string, field interface{}) (uint64, error) {
	return redis.Uint64(c.Conn.Do("HGET", c.serviceNamespace+key, field))
}

// HGetFloat64 get the float value of a hash field.
func (c *WingRedisConn) HGetFloat64(key string, field interface{}) (float64, error) {
	return redis.Float64(c.Conn.Do("HGET", c.serviceNamespace+key, field))
}

// HGetBytes get the bytes array of a hash field.
func (c *WingRedisConn) HGetBytes(key string, field interface{}) ([]byte, error) {
	return redis.Bytes(c.Conn.Do("HGET", c.serviceNamespace+key, field))
}

// HGetBool get the bool value of a hash field.
func (c *WingRedisConn) HGetBool(key string, field interface{}) (bool, error) {
	return redis.Bool(c.Conn.Do("HGET", c.serviceNamespace+key, field))
}
