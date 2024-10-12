package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/VladislavSCV/internal/config"
	"github.com/VladislavSCV/internal/database/postgres"
	"github.com/VladislavSCV/internal/model"
	"github.com/VladislavSCV/pkg"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func main() {
	go config.LoadEnv()
	db := postgres.NewUserPostgresHandlerDB(os.Getenv("CONN_TO_DB_PQ"))

	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Logger middleware will write the logs to gin.DefaultWriter stream
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.GET("/api/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/api/v1/get_users", func(c *gin.Context) {
		users, err := db.GetUsers()
		pkg.CError(err)
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	})

	r.POST("/api/v1/get_user_by_login", func(c *gin.Context) {
		var user model.User
		err := c.ShouldBindJSON(&user)
		pkg.CError(err)

		user, err = db.GetUserByLogin(user.Login)
		pkg.CError(err)

		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	})

	r.GET("/api/v1/get_user/:id", func(c *gin.Context) {
		strId := c.Param("id")
		id, err := strconv.Atoi(strId)
		pkg.CError(err)

		user, err := db.GetUserById(id)
		pkg.CError(err)

		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	})

	r.POST("/api/v1/create_user", func(c *gin.Context) {
		var user model.User
		err := c.ShouldBindJSON(&user)
		pkg.CError(err)

		err = db.CreateUser(user)
		pkg.CError(err)

		c.JSON(http.StatusOK, gin.H{
			"message": "create user",
		})
	})

	err := r.Run(":8000")
	if err != nil {
		return
	}
}
