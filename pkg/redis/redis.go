package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

func NewRedisConn(c RedisConfig) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     c.Host,
		Password: c.Password,
		DB:       c.Database,
	})

	if err := redisClient.Ping(context.TODO()).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	return redisClient, nil
}
