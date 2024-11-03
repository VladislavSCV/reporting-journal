package t_main

//
//import (
//	"errors"
//	"net/http"
//	"os"
//
//	"github.com/VladislavSCV/api/middleware"
//	"github.com/VladislavSCV/api/rest/handlers"
//	"github.com/VladislavSCV/internal/config"
//	"github.com/VladislavSCV/internal/users"
//	"github.com/VladislavSCV/pkg"
//	"github.com/gin-gonic/gin"
//)
//
//func main() {
//	config.LoadEnv()
//	dbp := users.NewUserPostgresHandlerDB(os.Getenv("CONN_TO_DB_PQ"))
//	dbr := users.NewUserHandlerRedis(os.Getenv("CONN_TO_REDIS"))
//	api := handlers.NewUserHandler(dbp, dbr)
//	r := gin.Default()
//
//	r.Use(gin.Logger())
//	r.Use(gin.Recovery())
//
//	r.GET("/api/v1/ping", func(c *gin.Context) {
//		c.JSON(http.StatusOK, gin.H{"message": "pong"})
//	})
//
//	r.GET("/api/v1/get_users", func(c *gin.Context) {
//		err := api.GetUsers(c)
//		if err != nil {
//			return
//		}
//	})
//
//	r.POST("/api/v1/login", func(c *gin.Context) {
//		user, err := api.GetUserByLogin(c)
//		if err != nil {
//			c.Status(http.StatusUnauthorized)
//			return
//		}
//		jwt, err := pkg.GenerateJWT(user.ID)
//		if err != nil {
//			c.Status(http.StatusUnauthorized)
//			return
//		}
//		c.JSON(http.StatusOK, gin.H{"token": jwt})
//	})
//
//	r.GET("/api/v1/get_user", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
//		userID, exists := c.Get("userID")
//		if !exists {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
//			return
//		}
//
//		// Здесь получить данные пользователя из Redis или PostgreSQL
//
//	})
//
//	r.GET("/api/v1/get_user/:id", func(c *gin.Context) {
//		err := api.GetUser(c)
//		if err != nil {
//			return
//		}
//	})
//
//	r.POST("/api/v1/registration", func(c *gin.Context) {
//		err := api.SignUp(c)
//		if err != nil {
//			return
//		}
//	})
//
//	r.PUT("/api/v1/update_user", func(c *gin.Context) {
//		err := api.UpdateUser(c)
//		if err != nil {
//			return
//		}
//	})
//
//	r.DELETE("/api/v1/delete_user/:id", func(c *gin.Context) {
//		err := api.DeleteUser(c)
//		if err != nil {
//			return
//		}
//	})
//
//	if err := r.Run(":8000"); err != nil {
//		pkg.LogWriteFileReturnError(errors.New("failed to start the server"))
//		return
//	}
//}
