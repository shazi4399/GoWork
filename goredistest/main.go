package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "39.106.49.188:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})
}

func main() {

	ctx := context.Background()
	err := rdb.Set(ctx, "key", "val", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key:", val)
}
