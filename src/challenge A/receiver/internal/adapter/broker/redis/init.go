package redis

import (
	"github.com/gomodule/redigo/redis"
	"strings"
)

// RedisQueue represents a Redis-backed queue for storing orders.
type RedisQueue struct {
	pool *redis.Pool
}

// NewRedisQueue creates a new RedisQueue with the given Redis connection pool.
func NewRedisQueue(pool *redis.Pool) *RedisQueue {
	return &RedisQueue{pool}
}

func InitRedis(dsn string) *RedisQueue {
	redisPool := &redis.Pool{
		MaxIdle:   5,
		MaxActive: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", strings.TrimSpace(dsn))
		},
	}
	redisQ := NewRedisQueue(redisPool)
	return redisQ
}
