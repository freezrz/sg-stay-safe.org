package pkg

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

type Redis struct {
	Client *redis.Client
}

func New(endpoint string) *Redis {
	log.Printf("Connecting to redis: %s/%d", endpoint, 0)
	redisClient := &Redis{redis.NewClient(&redis.Options{
		Addr:         endpoint,
		Password:     "",
		DB:           0,
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

func (r *Redis) Del(key string) error {
	defer func() {
		_ = r.Client.Close()
	}()

	iter := r.Client.Scan(0, key, 0).Iterator()
	log.Printf("del: %v", key)
	for iter.Next() {
		err := r.Client.Del(iter.Val()).Err()
		if err != nil {
			log.Println(err.Error())
			return err
		}
	}
	if err := iter.Err(); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (r *Redis) Close() error {
	err := r.Client.Close()
	if err != nil {
		return err
	}
	return nil
}
