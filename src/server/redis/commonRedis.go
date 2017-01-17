package redis

import (
	"github.com/garyburd/redigo/redis"
	"goshop/src/server/config"
	"goshop/src/server/utils"
)

func Hset(key, field, value string) bool {
	if ok, err := redis.Bool(config.Redis.Do("HSET", key, field, value)); ok {
		utils.DelError(err)
		return true
	}
	return false
}

func Set(key, value string) bool {
	if ok, err := redis.Bool(config.Redis.Do("SET", key, value)); ok {
		utils.DelError(err)
		return true
	}
	return false
}

func Get(key string) string {
	if exists, err := redis.Bool(config.Redis.Do("EXISTS", key)); !exists {
		utils.DelError(err)
		return ""
	}
	if value, err := redis.String(config.Redis.Do("GET", key)); value != "" {
		utils.DelError(err)
		return value
	}
	return ""
}

func Expire(key string, expireTime int) bool {
	if ok, err := redis.Bool(config.Redis.Do("EXPIRE", key, expireTime)); ok {
		utils.DelError(err)
		return true
	}
	return false
}

func Delete(key string) bool {
	if ok, err := redis.Bool(config.Redis.Do("DEL", key)); ok {
		utils.DelError(err)
		return true
	}
	return false
}
