package db

import (
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-redis/redis"
	"time"
)

var (
	redisdb *redis.Client
)

func initRedis() (err error) {
	redisdb = redis.NewClient(
		&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		})
	_, err = redisdb.Ping().Result()
	return
}

func saveNotes(key string, value string) error {
	res, err := redisdb.Set(key, value, time.Minute*10).Result()
	if err != nil {
		return fmt.Errorf("save error, %s, %s", err.Error(), res)
	}
	return nil
}

func getNotes(key string) (string, error) {
	return redisdb.Get(key).Result()
}
