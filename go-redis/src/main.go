package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	// go-redisを使ったredisクライアント

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	count := 0
	for {
		select {
		case <-ticker.C:
			log.Printf("count=%d\n", count)
			count++
			res, err := getLock(ctx, rdb)
			if err != nil {
				log.Fatalf("Failed to get lock. err: %v", err)
				panic(err)
			}

			if !res {
				log.Printf("Failed to get lock. count=%d\n", count)
				continue
			}
			log.Printf("Success to get lock. count=%d, res:=%b", count, res)
		}
	}

}

func getLock(ctx context.Context, client *redis.Client) (bool, error) {
	// setnxでロックを取得する
	res, err := client.SetNX(ctx, "lock", "1", 10*time.Second).Result()
	if err != nil {
		return false, fmt.Errorf("Failed to execute redis SetNX. err: %v", err)
	}
	log.Printf("res: %v\n", res)
	return res, nil
}
