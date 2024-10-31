package groups

import (
	"context"
	"testing"

	"github.com/VladislavSCV/internal/model"
	"github.com/redis/go-redis/v9"
)

func setupTestRedis(t *testing.T) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		DB:       0,
		Password: "VLADISLAV@31415926",
		Addr:     "localhost:6379",
	})
	err := rdb.FlushAll(context.Background()).Err()
	if err != nil {
		t.Fatalf("Ошибка очистки Redis перед тестом: %v", err)
	}
	return rdb
}

func TestCacheGroups(t *testing.T) {
	rdb := setupTestRedis(t)
	defer rdb.Close()

	groupCache := NewGroupCache(rdb)
	groups := []*model.Group{{Name: "Группа 1"}, {Name: "Группа 2"}}

	err := groupCache.CacheGroups(groups)
	if err != nil {
		t.Fatalf("Ошибка кеширования групп: %v", err)
	}

	// Проверяем, что данные закешированы
	data, err := rdb.Get(context.Background(), "groups").Result()
	if err != nil {
		t.Fatalf("Ошибка получения данных из кеша: %v", err)
	}
	if data == "" {
		t.Errorf("Ожидались данные в кеше, но они отсутствуют")
	}
}
