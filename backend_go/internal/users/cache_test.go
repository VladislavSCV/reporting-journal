package users

import (
	"fmt"
	"testing"

	"github.com/VladislavSCV/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestUserHandlerRedis_Login(t *testing.T) {
	r := connToRedis("rediss://red-csbnmi9u0jms73ffg86g:Mjn001c9yc3dk2cD5wIZXZwcJJmbWytE@oregon-redis.render.com:6379")
	uhr := &userHandlerRedis{redisClient: r}
	defer r.Close()

	user := models.User{
		Name:     "t_main",
		RoleID:   1,
		GroupID:  1,
		Login:    "t_main",
		Password: "t_main",
	}

	err := uhr.Login(&user)
	assert.NoError(t, err)

	userKey := fmt.Sprintf("user:%d", user.ID)
	result, err := r.HGetAll(ctx, userKey).Result()
	assert.NoError(t, err)

	assert.Equal(t, user.Name, result["name"])
	assert.Equal(t, fmt.Sprint(user.RoleID), result["role_id"])
	assert.Equal(t, fmt.Sprint(user.GroupID), result["group_id"])
	assert.Equal(t, user.Login, result["login"])
	assert.Equal(t, user.Password, result["password"])
}

func TestUserHandlerRedis_Logout(t *testing.T) {
	r := connToRedis("rediss://red-csbnmi9u0jms73ffg86g:Mjn001c9yc3dk2cD5wIZXZwcJJmbWytE@oregon-redis.render.com:6379")
	uhr := &userHandlerRedis{redisClient: r}
	defer r.Close()

	user := models.User{
		Name:     "t_main",
		RoleID:   1,
		GroupID:  1,
		Login:    "t_main",
		Password: "t_main",
	}

	err := uhr.Login(&user)
	assert.NoError(t, err)

	err = uhr.Logout(user.ID)
	assert.NoError(t, err)

	userKey := fmt.Sprintf("user:%d", user.ID)
	_, err = r.HGetAll(ctx, userKey).Result()
	if err != nil {
		t.Errorf("Expected redis.Nil, but got %v", err)
	}
}

func TestUserHandlerRedis_GetUser(t *testing.T) {
	r := connToRedis("rediss://red-csbnmi9u0jms73ffg86g:Mjn001c9yc3dk2cD5wIZXZwcJJmbWytE@oregon-redis.render.com:6379")
	uhr := &userHandlerRedis{redisClient: r}
	defer r.Close()

	user := models.User{
		Name:     "t_main",
		RoleID:   1,
		GroupID:  1,
		Login:    "t_main",
		Password: "t_main",
	}

	err := uhr.Login(&user)
	assert.NoError(t, err)

	userData, err := uhr.GetUserById(user.ID)
	assert.NoError(t, err)

	assert.Equal(t, user.ID, userData.ID)
	assert.Equal(t, user.Name, userData.Name)
	assert.Equal(t, user.RoleID, userData.RoleID)
	assert.Equal(t, user.GroupID, userData.GroupID)
	assert.Equal(t, user.Login, userData.Login)
	assert.Equal(t, user.Password, userData.Password)
}

func TestUserHandlerRedis_UpdateUser(t *testing.T) {
	r := connToRedis("rediss://red-csbnmi9u0jms73ffg86g:Mjn001c9yc3dk2cD5wIZXZwcJJmbWytE@oregon-redis.render.com:6379")
	uhr := &userHandlerRedis{redisClient: r}
	defer r.Close()

	user := models.User{
		ID:       1, // Ensure ID is set
		Name:     "test12",
		RoleID:   1,
		GroupID:  1,
		Login:    "t_main",
		Password: "t_main",
	}

	err := uhr.Login(&user)
	assert.NoError(t, err)

	updates := map[string]string{
		"name":     "test_updated", // New value for name
		"role_id":  "2",            // Updated value for role_id
		"group_id": "2",            // Updated value for group_id
	}

	err = uhr.UpdateUser(fmt.Sprint(user.ID), updates)
	assert.NoError(t, err)

	userData, err := uhr.GetUserById(user.ID)
	assert.NoError(t, err)

	assert.Equal(t, user.ID, userData.ID)
	assert.Equal(t, updates["name"], userData.Name) // Use the updated name from the map
	assert.Equal(t, 2, userData.RoleID)             // Updated role_id
	assert.Equal(t, 2, userData.GroupID)            // Updated group_id
	// Updated group_id
	assert.Equal(t, user.Login, userData.Login)
	assert.Equal(t, user.Password, userData.Password)
}

func TestUserHandlerRedis_DeleteUser(t *testing.T) {
	r := connToRedis("rediss://red-csbnmi9u0jms73ffg86g:Mjn001c9yc3dk2cD5wIZXZwcJJmbWytE@oregon-redis.render.com:6379")
	uhr := &userHandlerRedis{redisClient: r}
	defer r.Close()

	user := models.User{
		ID:       1, // Ensure ID is set for the t_main
		Name:     "t_main",
		RoleID:   1,
		GroupID:  1,
		Login:    "t_main",
		Password: "t_main",
	}

	err := uhr.Login(&user)
	assert.NoError(t, err)

	// Delete the user
	err = uhr.DeleteUser(user.ID)
	assert.NoError(t, err)

	// Verify the user was deleted
	userKey := fmt.Sprintf("user:%d", user.ID)
	_, err = r.HGetAll(ctx, userKey).Result()
	assert.Equal(t, nil, err) // Check for redis.Nil
}
