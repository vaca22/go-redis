package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	rdb *redis.Client
)

func V8Example() {
	ctx := context.Background()

	err := rdb.Set(ctx, "key", "valudde", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//_, err := rdb.Ping(ctx).Result()
	V8Example()

}
