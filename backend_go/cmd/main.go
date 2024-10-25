package main

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/VladislavSCV/internal/cache"
	"github.com/VladislavSCV/internal/config"
	"github.com/VladislavSCV/internal/database/postgres"
	"github.com/VladislavSCV/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	dbp := postgres.NewUserPostgresHandlerDB(os.Getenv("CONN_TO_DB_PQ"))
	dbr := cache.NewUserHandlerRedis(os.Getenv("CONN_TO_REDIS"))
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/api/v1/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/api/v1/get_users", func(c *gin.Context) {
		users, err := dbp.GetUsers()
		if err != nil {
			pkg.LogWriteFileReturnError(errors.New("failed to retrieve users from the database"))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"users": users})
	})

	r.GET("/api/v1/get_user_by_login/:login", func(c *gin.Context) {
		strLogin := c.Param("login")
		user, err := dbp.GetUserByLogin(strLogin)

		if err != nil {
			pkg.LogWriteFileReturnError(errors.New("failed to retrieve user by login"))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": user})
	})

	r.GET("/api/v1/get_user/:id", func(c *gin.Context) {
		strId := c.Param("id")
		id, err := strconv.Atoi(strId)
		if err != nil {
			pkg.LogWriteFileReturnError(errors.New("invalid user ID format"))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		user, err := dbr.GetUserById(id)
		if err != nil {
			pkg.LogWriteFileReturnError(errors.New("failed to retrieve user from Redis"))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
			return
		}
		if user.Name == "" {
			user, err = dbp.GetUserById(id)
			if err != nil {
				pkg.LogWriteFileReturnError(errors.New("failed to retrieve user from PostgreSQL"))
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
				return
			}
			user.ID = id
		}

		c.JSON(http.StatusOK, gin.H{"user": user})
	})

	r.POST("/api/v1/create_user", func(c *gin.Context) {
		user, err := pkg.ParseUserRequest(c)
		if err != nil {
			pkg.LogWriteFileReturnError(errors.New("invalid request payload for creating user"))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		err = dbr.Login(user)
		if err != nil {
			err = dbp.CreateUser(user)
			if err != nil {
				pkg.LogWriteFileReturnError(errors.New("failed to create user in PostgreSQL"))
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	})

	r.PUT("/api/v1/update_user", func(c *gin.Context) {
		userUpdates := make(map[string]string)

		// Parse updates from JSON
		if err := c.ShouldBindJSON(&userUpdates); err != nil {
			pkg.LogWriteFileReturnError(errors.New("invalid request payload for updating user"))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		// Ensure "id" field exists
		userID, exists := userUpdates["id"]
		if !exists {
			pkg.LogWriteFileReturnError(errors.New("missing 'id' field in request body"))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user ID"})
			return
		}

		err := dbr.UpdateUser(userID, userUpdates)
		if err != nil {
			err = dbp.UpdateUser(userID, userUpdates)
			if err != nil {
				pkg.LogWriteFileReturnError(errors.New("failed to update user in PostgreSQL"))
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
	})

	r.DELETE("/api/v1/delete_user/:id", func(c *gin.Context) {
		strId := c.Param("id")
		id, err := strconv.Atoi(strId)
		if err != nil {
			pkg.LogWriteFileReturnError(errors.New("invalid user ID format for deletion"))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		err = dbr.DeleteUser(id)
		if err != nil {
			err = dbp.DeleteUser(id)
			if err != nil {
				pkg.LogWriteFileReturnError(errors.New("failed to delete user from PostgreSQL"))
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	})

	if err := r.Run(":8000"); err != nil {
		pkg.LogWriteFileReturnError(errors.New("failed to start the server"))
		return
	}
}
