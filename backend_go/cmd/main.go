package main

import (
	"context"
	"github.com/VladislavSCV/api/middleware"
	"github.com/VladislavSCV/internal/config"
	"github.com/VladislavSCV/internal/note"
	"github.com/VladislavSCV/internal/subjects"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/VladislavSCV/api/rest/handlers"
	"github.com/VladislavSCV/internal/groups"
	"github.com/VladislavSCV/internal/role"
	"github.com/VladislavSCV/internal/schedules"
	"github.com/VladislavSCV/internal/users"
	"github.com/gin-gonic/gin"
)

type ApiHandlers struct {
	UserApi     users.UserAPIRepository
	RoleApi     role.RoleApiRepository
	GroupApi    groups.GroupApiRepository
	NoteApi     note.NoteApiRepository
	ScheduleApi schedules.ScheduleApiRepository
	SubjectApi  subjects.SubjectApiRepository
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
	r.Use(middleware.CORSMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(func(c *gin.Context) {
		log.Printf("Request headers: %v", c.Request.Header)
		c.Next()
	})

	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/", errorHandler(api.UserApi.GetUserByToken))
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
		groupRoutes.GET("/:id", errorHandler(api.GroupApi.GetGroupByID))
		// TODO настроить возврат id группы
		groupRoutes.POST("/", errorHandler(api.GroupApi.CreateGroup))
		// TODO настроить возврат id группы
		groupRoutes.PUT("/:id", errorHandler(api.GroupApi.UpdateGroup))
		groupRoutes.DELETE("/:id", errorHandler(api.GroupApi.DeleteGroup))
	}

	notesRoutes := r.Group("/api/note")
	{
		notesRoutes.GET("/", errorHandler(api.NoteApi.GetNotes))
		notesRoutes.GET("/:id", errorHandler(api.NoteApi.GetNote))
		notesRoutes.POST("/", errorHandler(api.NoteApi.CreateNote))
		notesRoutes.PUT("/:id", errorHandler(api.NoteApi.UpdateNote))
		notesRoutes.DELETE("/:id", errorHandler(api.NoteApi.DeleteNote))
	}

	scheduleRoutes := r.Group("/api/schedule")
	{
		scheduleRoutes.GET("/", errorHandler(api.ScheduleApi.GetSchedules))
		scheduleRoutes.GET("/:id", errorHandler(api.ScheduleApi.GetSchedule))
		scheduleRoutes.POST("/", errorHandler(api.ScheduleApi.CreateSchedule))
		scheduleRoutes.PUT("/:id", errorHandler(api.ScheduleApi.UpdateSchedule))
		scheduleRoutes.DELETE("/:id", errorHandler(api.ScheduleApi.DeleteSchedule))
	}

	subjectRoutes := r.Group("/api/subject")
	{
		subjectRoutes.GET("/", errorHandler(api.SubjectApi.GetSubjects))
		subjectRoutes.GET("/:id", errorHandler(api.SubjectApi.GetSubjectById))
		subjectRoutes.POST("/", errorHandler(api.SubjectApi.CreateSubject))
		subjectRoutes.PUT("/:id", errorHandler(api.SubjectApi.UpdateSubject))
		subjectRoutes.DELETE("/:id", errorHandler(api.SubjectApi.DeleteSubject))
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

	dbpn := note.NewNotePostgresHandlerDB(connToDb)
	apiNotes := handlers.NewNoteHandler(dbpn)

	dbps := schedules.NewUserPostgresHandlerDB(connToDb)
	apiSchedules := handlers.NewScheduleHandler(dbps)

	dbpsu := subjects.NewSubjectPostgresHandlerDB(connToDb)
	apiSubjects := handlers.NewSubjectHandler(dbpsu)

	api := ApiHandlers{UserApi: apiUsers, RoleApi: apiRoles, GroupApi: apiGroups, NoteApi: apiNotes, ScheduleApi: apiSchedules, SubjectApi: apiSubjects}
	router := SetupRouter(api)
	srv := &http.Server{
		Addr:    "localhost:8000",
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
