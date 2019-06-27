package main

import (
"fmt"
"github.com/go-redis/redis"
"time"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		})
	defer client.Close()

	pubsub := client.Subscribe("mychannel1")
	defer pubsub.Close()

	err := client.Publish("mychannel1", "hello world!").Err()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Published Message time %v\n", time.Now().UnixNano())
}
