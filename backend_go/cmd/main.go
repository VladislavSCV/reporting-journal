package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/VladislavSCV/api/middleware"
	"github.com/VladislavSCV/api/rest/handlers"
	"github.com/VladislavSCV/internal/config"
	"github.com/VladislavSCV/internal/users"
	"github.com/gin-gonic/gin"
)

//type allHandlers struct {
//	UserApi users.UserAPIRepository
//	RoleApi
//}

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func SetupRouter(api users.UserAPIRepository) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//protected := r.Group("/")

	//protected.Use(middleware.AuthMiddleware())
	//{
	//
	//}

	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/registration", errorHandler(api.SignUp))
		authRoutes.POST("/login", errorHandler(api.Login))
		authRoutes.POST("/verify", errorHandler(api.VerifyToken))
	}

	userRoutes := r.Group("/api/user")
	{
		userRoutes.GET("/", errorHandler(api.GetUsers))
		userRoutes.GET("/:id", errorHandler(api.GetUser))
		userRoutes.PUT("/:id", errorHandler(api.UpdateUser))
		userRoutes.DELETE("/:id", errorHandler(api.DeleteUser))
	}

	//roleRoutes := r.Group("/api/role")
	//{
	//	roleRoutes.GET("/", errorHandler(api.GetRoles))
	//	roleRoutes.GET("/:id", errorHandler(api.GetRole))
	//	roleRoutes.POST("/", errorHandler(api.CreateRole))
	//	roleRoutes.PUT("/:id", errorHandler(api.UpdateRole))
	//	roleRoutes.DELETE("/:id", errorHandler(api.DeleteRole))
	//}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
	})

	return r
}

func main() {
	config.LoadEnv()
	connToDb := os.Getenv("CONN_TO_DB_PQ")
	if connToDb == "" {
		log.Fatal("CONN_TO_DB_PQ environment variable is not set")
	}

	dbp := users.NewUserPostgresHandlerDB(connToDb)
	dbr := users.NewUserHandlerRedis(os.Getenv("CONN_TO_REDIS"))
	api := handlers.NewUserHandler(dbp, dbr)

	router := SetupRouter(api)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}

func errorHandler(handler func(c *gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := handler(c); err != nil {
			var statusCode int
			var message string

			// Type assertions to determine the type of error
			switch e := err.(type) {
			case *NotFoundError:
				statusCode = http.StatusNotFound
				message = e.Error()
			case *ValidationError:
				statusCode = http.StatusBadRequest
				message = e.Error()
			default:
				statusCode = http.StatusInternalServerError
				message = "Internal Server Error"
			}

			// Respond with the appropriate error message
			c.JSON(statusCode, gin.H{"error": message})
		}
	}
}
