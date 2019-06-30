package configs

import (
	"github.com/go-redis/redis"
)

// RedisOptions returns the default redis options to use for this project
func RedisOptions() *redis.Options {
	return &redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}
}
