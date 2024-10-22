package redis

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/VladislavSCV/internal/model"
	"github.com/VladislavSCV/pkg"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type UserHandlerRedisInterface interface {
	SetUser(user model.User) error
	GetUser(id int) (model.User, error)
	UpdateUser(id string, updates map[string]string) error
	DeleteUser(id int) error
}

type userHandlerRedis struct {
	redisClient *redis.Client
}

func (uhr *userHandlerRedis) SetUser(user model.User) error {
	userId := fmt.Sprintf("user:%d", user.ID)
	// Expiration - время жизни ключа в Redis, 10000 - 10 секунд
	//uhr.redisClient.Set(ctx, userId, user, 2880*60000) // 0 - ключ не будет истекать

	err := uhr.redisClient.HSet(ctx, userId, "name", user.Name, "role_id", user.RoleID, "group_id", user.GroupID, "login", user.Login, "password", user.Password).Err()
	if err != nil {
		return err
	}
	return nil
}

func (uhr *userHandlerRedis) GetUser(id int) (model.User, error) {
	userId := fmt.Sprintf("user:%d", id)
	user, err := uhr.redisClient.HGetAll(ctx, userId).Result()
	if err != nil {
		return model.User{}, err
	}

	roleId, err := strconv.Atoi(user["role_id"])
	if err != nil {
		pkg.LogWriteFileReturnError(err)
	}

	groupId, err := strconv.Atoi(user["group_id"])
	if err != nil {
		pkg.LogWriteFileReturnError(err)
	}

	userData := model.User{
		ID:       id,
		Name:     user["name"],
		RoleID:   roleId,
		GroupID:  groupId,
		Login:    user["login"],
		Password: user["password"],
	}

	return userData, nil
}

func (uhr *userHandlerRedis) UpdateUser(id string, updates map[string]string) error {
	err := uhr.redisClient.HSet(ctx, id, updates).Err()
	if err != nil {
		return err
	}

	return nil
}

func (uhr *userHandlerRedis) DeleteUser(id int) error {
	userId := fmt.Sprintf("user:%d", id)
	err := uhr.redisClient.Del(ctx, userId).Err()
	if err != nil {
		return err
	}
	return nil
}

func connToRedis(connStr string) *redis.Client {
	opt, err := redis.ParseURL(os.Getenv("CONN_TO_REDIS"))
	if err != nil {
		pkg.LogWriteFileReturnError(err)
		return nil
	}

	return redis.NewClient(opt)
}

func checkConn(client *redis.Client) error {
	if err := client.Ping(ctx).Err(); err != nil {
		return err
	}
	return nil
}

func NewUserHandlerRedis(connStr string) UserHandlerRedisInterface {
	r := connToRedis(connStr)
	return &userHandlerRedis{redisClient: r}
}
