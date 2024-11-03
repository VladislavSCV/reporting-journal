package groups

import (
	"context"
	"encoding/json"
	"time"

	"github.com/VladislavSCV/internal/models"
	"github.com/redis/go-redis/v9"
)

type groupCache struct {
	redisClient *redis.Client
}

func NewGroupCache(redisClient *redis.Client) GroupRedisRepository {
	return &groupCache{redisClient: redisClient}
}

// CacheGroups сохраняет список групп в Redis с временем жизни кеша
func (gc *groupCache) CacheGroups(groups []*models.Group) error {
	data, err := json.Marshal(groups)
	if err != nil {
		return err
	}

	return gc.redisClient.Set(context.Background(), "groups", data, 10*time.Minute).Err()
}

// GetCachedGroups получает список групп из кеша Redis
func (gc *groupCache) GetCachedGroups() ([]*models.Group, error) {
	data, err := gc.redisClient.Get(context.Background(), "groups").Result()
	if err == redis.Nil {
		return nil, nil // Кеш пуст
	} else if err != nil {
		return nil, err
	}

	var groups []*models.Group
	if err := json.Unmarshal([]byte(data), &groups); err != nil {
		return nil, err
	}

	return groups, nil
}
