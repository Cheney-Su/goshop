package config

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"github.com/spf13/viper"
	"log"
)

func SetUpRedis() {
	addr := viper.GetString("redis.addr")
	password := viper.GetString("redis.password")
	database := viper.GetInt("redis.database")
	maxIdle := viper.GetInt("redis.max-idle")
	maxActive := viper.GetInt("redis.max-active")
	idleTimeout := time.Duration(viper.GetInt("redis.idle-timeout")) * time.Second

	pool := redis.Pool{
		MaxActive:maxActive,
		MaxIdle:maxIdle,
		IdleTimeout:idleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			if c.Do("SELECT", database); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	log.Print("创建redis连接池成功...")
	Redis = pool.Get()
}

var (
	Redis redis.Conn
)