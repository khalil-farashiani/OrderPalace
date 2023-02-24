package redis

import (
	"github.com/gomodule/redigo/redis"
)

// RedisQueue represents a Redis-backed queue for storing orders.
type RedisQueue struct {
	conn redis.Conn
}

// NewRedisQueue creates a new RedisQueue with the given Redis connection pool.
func NewRedisQueue(conn redis.Conn) *RedisQueue {
	return &RedisQueue{conn: conn}
}

func InitRedis(dsn string) *RedisQueue {
	// Connect to Redis
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	redisQ := NewRedisQueue(conn)
	return redisQ
}
