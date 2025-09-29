package database

import "github.com/redis/go-redis/v9"

func NewRedisClient(dsn string) (*redis.Client, error) {
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(opt)
	return rdb, nil
}
