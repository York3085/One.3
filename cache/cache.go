package cache

import (
	"One.3/dao"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func GetUserInformation(username string) (string, error) {
	cacheClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	cachedData, err := cacheClient.Get(ctx, username).Result()
	if err == nil {
		return cachedData, nil
	}

	password := dao.SelectPasswordFromUsername(username)
	if password == "" {
		return "", fmt.Errorf("user not found")
	}
	return password, nil
}
