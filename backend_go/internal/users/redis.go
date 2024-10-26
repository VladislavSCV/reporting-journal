package users

import (
	"context"
	"fmt"
	"strconv"

	"github.com/VladislavSCV/internal/model"
	"github.com/VladislavSCV/pkg"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type userHandlerRedis struct {
	redisClient *redis.Client
}

// Login добавляет нового пользователя в Redis
//
//	@param user model.User - пользователь, который будет добавлен
//
//	@return error - ошибка, если она возникла
func (uhr *userHandlerRedis) Login(user *model.User) error {
	userKey := fmt.Sprintf("user:%d", user.ID)
	// Expiration - время жизни ключа в Redis, 10000 - 10 секунд
	//uhr.redisClient.Set(ctx, userId, user, 0) // 0 - ключ не будет истекать

	err := uhr.redisClient.HSet(ctx, userKey, "name", &user.Name, "role_id", &user.RoleID, "group_id", &user.GroupID, "login", &user.Login, "password", &user.Password).Err()
	if err != nil {
		return err
	}
	return nil
}

// Logout удаляет существующего пользователя
//
//	@param userId int - ID пользователя, который будет удален
//
//	@return error - ошибка, если она возникла
func (uhr *userHandlerRedis) Logout(id int) error {
	userKey := fmt.Sprintf("user:%d", id)
	_, err := uhr.redisClient.Del(ctx, userKey).Result()
	return err
}

// GetUserById возвращает пользователя по его ID
//
//	@param id int - ID пользователя
//
//	@return model.User - пользователь
//	@return error - ошибка, если она возникла
func (uhr *userHandlerRedis) GetUserById(id int) (model.User, error) {
	userKey := fmt.Sprintf("user:%d", id)
	user, err := uhr.redisClient.HGetAll(ctx, userKey).Result()
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

// UpdateUser обновляет существующего пользователя
//
//	@param id string - ID пользователя, который будет обновлен
//	@param updates map[string]string - поля, которые будут обновлены
//
//	@return error - ошибка, если она возникла
func (uhr *userHandlerRedis) UpdateUser(id string, updates map[string]string) error {
	// Формируем ключ для пользователя в Redis
	userKey := fmt.Sprintf("user:%s", id)

	// Получаем текущие данные пользователя
	currentData, err := uhr.redisClient.HGetAll(ctx, userKey).Result()
	if err != nil {
		return err
	}

	// Обновляем поля в currentData только с новыми значениями из updates
	for key, value := range updates {
		currentData[key] = value
	}

	// Преобразуем обновленные данные обратно в Redis
	err = uhr.redisClient.HMSet(ctx, userKey, currentData).Err()
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser удаляет существующего пользователя
//
//	@param id int - ID пользователя, который будет удален
//
//	@return error - ошибка, если она возникла
func (uhr *userHandlerRedis) DeleteUser(id int) error {
	userId := fmt.Sprintf("user:%d", id)
	err := uhr.redisClient.Del(ctx, userId).Err()
	if err != nil {
		return err
	}
	return nil
}

// connToRedis возвращает клиент Redis, созданный на основе строки подключения.
//
//	@param connStr string - строка подключения к Redis
//
//	@return *redis.Client - клиент Redis
//	@return error - ошибка, если она возникла
func connToRedis(connStr string) *redis.Client {
	opt, err := redis.ParseURL(connStr)
	if err != nil {
		pkg.LogWriteFileReturnError(err)
		return nil
	}

	return redis.NewClient(opt)
}

// checkConn проверяет соединение с Redis
//
//	@param client *redis.Client - клиент Redis
//
//	@return error - ошибка, если она возникла
func checkConnRedis(client *redis.Client) error {
	if err := client.Ping(ctx).Err(); err != nil {
		return err
	}
	return nil
}

// NewUserHandlerRedis возвращает User
func NewUserHandlerRedis(connStr string) UserRedisRepository {
	r := connToRedis(connStr)

	err := checkConnRedis(r)
	if err != nil {
		pkg.LogWriteFileReturnError(err)
	}
	return &userHandlerRedis{redisClient: r}
}
