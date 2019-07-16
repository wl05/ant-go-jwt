package utils

import (
	"fmt"
	"github.com/go-redis/redis"
)

var RClient *redis.Client
// 连接redis
func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123123", // no password set
		DB:       0,        // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
	RClient = client
	return client
}
