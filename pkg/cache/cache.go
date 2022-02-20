package cache

import (
	"errors"
	"github.com/go-redis/redis"
	"log"
	"time"
)

type Redis struct {
	Client *redis.Client
}

func New(endpoint string) *Redis {
	log.Printf("connect to redis: %s/%d", endpoint, 0)
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
	log.Printf("get key: %s", key)
	results := r.Client.Get(key)
	if results != nil {
		value, err := r.Client.Get(key).Result()
		if err != nil {
			log.Println(err.Error()) // Redis `GET key` command. It returns redis.Nil error when key does not exist
			return "", nil
		}
		log.Printf("value: %s", value)
		return value, nil
	} else {
		return "", errors.New("fail to get cache")
	}
}

func (r *Redis) Set(key string, value interface{}, ttl int) error {
	log.Printf("Set Key:%s, Value:%v", key, value)
	err := r.Client.Set(key, value, time.Duration(ttl)*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Incr(key string, ttl ...int) error {
	log.Printf("Incr Key:%s", key)
	val := r.Client.Incr(key).Val()
	if val == 1 {
		if len(ttl) == 1 {
			r.Client.Expire(key, time.Minute*time.Duration(ttl[0]))
		}
	}
	return nil
}

func (r *Redis) Del(key string) error {
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
