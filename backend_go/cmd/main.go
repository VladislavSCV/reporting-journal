package main

import (
	"context"
	"github.com/VladislavSCV/internal/config"
	"github.com/gin-contrib/cors"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/VladislavSCV/api/rest/handlers"
	"github.com/VladislavSCV/internal/groups"
	"github.com/VladislavSCV/internal/role"
	"github.com/VladislavSCV/internal/users"
	"github.com/gin-gonic/gin"
)

type ApiHandlers struct {
	UserApi  users.UserAPIRepository
	RoleApi  role.RoleApiRepository
	GroupApi groups.GroupApiRepository
}

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

func SetupRouter(api ApiHandlers) *gin.Engine {
	r := gin.Default()
	// Настроим CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                   // Разрешаем запросы только с этого адреса
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // Разрешаем эти методы
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Разрешаем эти заголовки
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/registration", errorHandler(api.UserApi.SignUp))
		authRoutes.POST("/login", errorHandler(api.UserApi.Login))
		authRoutes.POST("/verify", errorHandler(api.UserApi.VerifyToken))
	}

	userRoutes := r.Group("/api/user")
	{
		userRoutes.GET("/", errorHandler(api.UserApi.GetUsers))
		userRoutes.GET("/:id", errorHandler(api.UserApi.GetUser))
		userRoutes.PUT("/:id", errorHandler(api.UserApi.UpdateUser))
		userRoutes.DELETE("/:id", errorHandler(api.UserApi.DeleteUser))
	}

	roleRoutes := r.Group("/api/role")
	{
		roleRoutes.GET("/", errorHandler(api.RoleApi.GetRoles))
		roleRoutes.GET("/:id", errorHandler(api.RoleApi.GetRole))
		roleRoutes.POST("/", errorHandler(api.RoleApi.CreateRole))
		roleRoutes.DELETE("/:id", errorHandler(api.RoleApi.DeleteRole))
	}

	groupRoutes := r.Group("/api/group")
	{
		groupRoutes.GET("/", errorHandler(api.GroupApi.GetGroups))
		groupRoutes.POST("/", errorHandler(api.GroupApi.CreateGroup))
		groupRoutes.PUT("/:id", errorHandler(api.GroupApi.UpdateGroup))
		groupRoutes.DELETE("/:id", errorHandler(api.GroupApi.DeleteGroup))
	}

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

	dbpu := users.NewUserPostgresHandlerDB(connToDb)
	dbru := users.NewUserHandlerRedis(os.Getenv("CONN_TO_REDIS"))
	apiUsers := handlers.NewUserHandler(dbpu, dbru)

	dbpr := role.NewRolePostgresHandler(connToDb)
	apiRoles := handlers.NewRoleHandler(dbpr)

	dbpg := groups.NewGroupPostgresRepository(connToDb)
	apiGroups := handlers.NewGroupHandler(dbpg)

	api := ApiHandlers{UserApi: apiUsers, RoleApi: apiRoles, GroupApi: apiGroups}
	router := SetupRouter(api)
	srv := &http.Server{
		Addr:    ":8000",
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
