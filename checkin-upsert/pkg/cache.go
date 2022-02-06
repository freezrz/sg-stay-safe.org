package pkg

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

type Redis struct {
	Client *redis.Client
}

type RedisConf struct {
	Endpoint string
	Db       int
	Password string
}

var config = RedisConf{
	Endpoint: "try-it-cache-delete-later.vekkvr.0001.apse1.cache.amazonaws.com:6379",
	Db:       0,
	Password: "",
}

func New() *Redis {
	log.Printf("Connecting to redis: %s/%d", config.Endpoint, config.Db)
	redisClient := &Redis{redis.NewClient(&redis.Options{
		Addr:         config.Endpoint,
		Password:     config.Password,
		DB:           config.Db,
		ReadTimeout:  10 * time.Second,
		DialTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})}
	return redisClient
}

func (r *Redis) Get(key string) (string, error) {
	log.Printf("Get Key: %s", key)
	value, err := r.Client.Get(key).Result()
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Fail to Get Key: %s. Retrying", key)
			r.Client.Get(key).Result()
		}
	}()
	return value, nil
}

func (r *Redis) Set(key string, value interface{}, ttl int) error {
	log.Printf("Set Key:%s, Value:%v", key, value)
	err := r.Client.Set(key, value, time.Duration(ttl)*time.Minute).Err()
	if err != nil {
		return err
	}
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Fail to Set Key:%s, Value:%v. Retrying", key, value)
			r.Client.Set(key, value, time.Duration(ttl)*time.Minute).Err()
		}
	}()
	return nil
}

func (r *Redis) Close() error {
	err := r.Client.Close()
	if err != nil {
		return err
	}
	return nil
}
