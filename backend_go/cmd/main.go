package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/VladislavSCV/internal/config"
	"github.com/VladislavSCV/internal/database/postgres"
	"github.com/VladislavSCV/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db := postgres.NewUserPostgresHandlerDB(os.Getenv("CONN_TO_DB_PQ"))

	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/api/v1/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/api/v1/get_users", func(c *gin.Context) {
		users, err := db.GetUsers()
		if err != nil {
			pkg.HandleHTTPError(c, http.StatusInternalServerError, "Failed to get users", err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"users": users})
	})

	r.GET("/api/v1/get_user_by_login/:login", func(c *gin.Context) {
		strLogin := c.Param("login")

		user, err := db.GetUserByLogin(strLogin)
		if err != nil {
			pkg.HandleHTTPError(c, http.StatusInternalServerError, "Failed to find user", err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": user})
	})

	r.GET("/api/v1/get_user/:id", func(c *gin.Context) {
		strId := c.Param("id")
		id, err := strconv.Atoi(strId)
		if err != nil {
			pkg.HandleHTTPError(c, http.StatusBadRequest, "Invalid user ID", err)
			return
		}

		user, err := db.GetUserById(id)
		if err != nil {
			pkg.HandleHTTPError(c, http.StatusInternalServerError, "Failed to find user", err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": user})
	})

	r.POST("/api/v1/create_user", func(c *gin.Context) {
		user, err := pkg.ParseUserRequest(c)
		if err != nil {
			pkg.HandleHTTPError(c, http.StatusBadRequest, "Invalid request payload", err)
			return
		}

		err = db.CreateUser(*user)
		if err != nil {
			pkg.HandleHTTPError(c, http.StatusInternalServerError, "Failed to create user", err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	})

	//r.POST("/api/v1/update_user", func(c *gin.Context) {
	//	user, err := pkg.ParseUserRequest(c)
	//	if err != nil {
	//		pkg.HandleHTTPError(c, http.StatusBadRequest, "Invalid request payload", err)
	//		return
	//	}
	//
	//	db.UpdateUser(user.ID, )
	//})

	if err := r.Run(":8000"); err != nil {
		return
	}
}
