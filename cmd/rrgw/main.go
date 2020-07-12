package main

import (
	"context"
	"fmt"
	"os"

	"github.com/amfl/redis_rest_gateway/pkg"
	"github.com/go-redis/redis/v8"
)

func main() {
	listenInterface := os.Args[1]
	redisAddr := os.Args[2]

	// Create a redis client we can use for publishing
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Create a gateway instance, which will act as a webserver
	gateway := redis_rest_gateway.Gateway{
		Client:  client,
		Context: context.Background(),
	}

	fmt.Println("Listening on", listenInterface)
	gateway.Listen(listenInterface)
}
