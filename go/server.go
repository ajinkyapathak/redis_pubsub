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

	fmt.Println("Subscribed to channel mychannel1")
	for{
		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			panic(err)
		}
		fmt.Printf("received time %v \n", time.Now().UnixNano())
		fmt.Println(msg.Channel, msg.Payload)
		time.Sleep(2000 * time.Millisecond)
	}
}
