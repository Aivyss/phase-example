package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

const KeywordExpirationTime time.Duration = 30 * time.Minute

type CacheRepository interface {
	Insert(key string) error
	Exists(key string) bool
}

type keywordRedisClient struct {
	client         *redis.ClusterClient
	expirationTime time.Duration
}

func (c *keywordRedisClient) Exists(key string) bool {
	val, err := c.client.Exists(context.Background(), key).Result()
	if err != nil {
		return false
	}

	return val > 0
}

// NewUserKeywordRepository UserKeywordRepository interfaceの値を生成する（Redis）
func NewUserKeywordRepository(client *redis.ClusterClient) CacheRepository {
	return &keywordRedisClient{
		client:         client,
		expirationTime: KeywordExpirationTime,
	}
}

func (c *keywordRedisClient) Insert(key string) error {
	return c.client.Set(context.Background(), key, true, c.expirationTime).Err()
}
