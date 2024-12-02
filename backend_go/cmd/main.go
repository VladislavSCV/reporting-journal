package main

import (
	"context"
	"github.com/VladislavSCV/api/middleware"
	"github.com/VladislavSCV/internal/attendance"
	"github.com/VladislavSCV/internal/config"
	"github.com/VladislavSCV/internal/models"
	"github.com/VladislavSCV/internal/note"
	"github.com/VladislavSCV/internal/subjects"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
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
	_ "net/http/pprof"
)

type ApiHandlers struct {
	UserApi     users.UserAPIRepository
	RoleApi     role.RoleApiRepository
	GroupApi    groups.GroupApiRepository
	NoteApi     note.NoteApiRepository
	ScheduleApi schedules.ScheduleApiRepository
	SubjectApi  subjects.SubjectApiRepository
	ElseApi     models.Else
}

const (
	StudentRoleID = 1
	TeacherRoleID = 2
	AdminRoleID   = 3
)

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
	r.Use(cors.Default())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(func(c *gin.Context) {
		log.Printf("Request headers: %v", c.Request.Header)
		c.Next()
	})

	// Маршруты для авторизации
	authRoutes := r.Group("/api/auth")
	{
		authRoutes.GET("/", api.UserApi.GetUserByToken)
		authRoutes.POST("/registration", api.UserApi.SignUp)
		authRoutes.POST("/login", api.UserApi.Login)
		authRoutes.POST("/verify", api.UserApi.VerifyToken)
	}

	// Маршруты для пользователей
	userRoutes := r.Group("/api/user")
	{
		userRoutes.GET("/", api.UserApi.GetUsers)            // возвращает список всех пользователей
		userRoutes.GET("/students", api.UserApi.GetStudents) // возвращает список всех студентов
		userRoutes.GET("/teachers", api.UserApi.GetTeachers) // возвращает список всех преподавателей
		userRoutes.GET("/:id", api.UserApi.GetUser)          // возвращает данные пользователя по ID
		userRoutes.PUT("/:id", api.UserApi.UpdateUser)       // обновляет данные пользователя
		userRoutes.DELETE("/:id", api.UserApi.DeleteUser)    // удаляет пользователя
	}

	// Маршруты для ролей
	roleRoutes := r.Group("/api/role")
	{
		roleRoutes.GET("/", errorHandler(api.RoleApi.GetRoles))         // возвращает список всех ролей
		roleRoutes.GET("/:id", errorHandler(api.RoleApi.GetRole))       // возвращает данные роли по ID
		roleRoutes.POST("/", errorHandler(api.RoleApi.CreateRole))      // создает новую роль
		roleRoutes.DELETE("/:id", errorHandler(api.RoleApi.DeleteRole)) // удаляет роль
	}

	// Маршруты для групп
	groupRoutes := r.Group("/api/group")
	{
		groupRoutes.GET("/", api.GroupApi.GetGroups)                  // возвращает список всех групп
		groupRoutes.GET("/schedule/:id", api.ScheduleApi.GetSchedule) // возвращает расписание для группы
		groupRoutes.GET("/:id", api.GroupApi.GetGroupByID)            // возвращает данные группы по ID
		groupRoutes.POST("/", api.GroupApi.CreateGroup)               // создает новую группу
		groupRoutes.PUT("/:id", api.GroupApi.UpdateGroup)             // обновляет данные группы
		groupRoutes.DELETE("/:id", api.GroupApi.DeleteGroup)          // удаляет группу
	}

	// Маршруты для заметок
	notesRoutes := r.Group("/api/note")
	{
		notesRoutes.GET("/", api.NoteApi.GetNotes)                              // возвращает список всех заметок
		notesRoutes.GET("/:id", api.NoteApi.GetNote)                            // возвращает данные заметки по ID
		notesRoutes.GET("/group/:id", api.NoteApi.GetGroupNote)                 // возвращает заметки для группы
		notesRoutes.GET("/curator/groups/:id", api.NoteApi.GetCuratorGroupNote) // возвращает заметки для группы для учителя
		notesRoutes.POST("/:id", api.NoteApi.CreateNote)                        // создает новую заметку
		notesRoutes.PUT("/:id", api.NoteApi.UpdateNote)                         // обновляет данные заметки
		notesRoutes.DELETE("/:id", api.NoteApi.DeleteNote)                      // удаляет заметку
	}

	// Маршруты для расписания
	scheduleRoutes := r.Group("/api/schedule")
	{
		scheduleRoutes.GET("/", api.ScheduleApi.GetSchedules)         // возвращает список всех расписаний
		scheduleRoutes.GET("/:id", api.ScheduleApi.GetSchedule)       // возвращает данные расписания по ID
		scheduleRoutes.POST("/", api.ScheduleApi.CreateSchedule)      // создает новое расписание
		scheduleRoutes.PUT("/:id", api.ScheduleApi.UpdateSchedule)    // обновляет данные расписания
		scheduleRoutes.DELETE("/:id", api.ScheduleApi.DeleteSchedule) // удаляет расписание
	}

	// Маршруты для предметов
	subjectRoutes := r.Group("/api/subject")
	{
		subjectRoutes.GET("/", api.SubjectApi.GetSubjects)         // возвращает список всех предметов
		subjectRoutes.GET("/:id", api.SubjectApi.GetSubjectById)   // возвращает данные предмета по ID
		subjectRoutes.POST("/", api.SubjectApi.CreateSubject)      // создает новый предмет
		subjectRoutes.PUT("/:id", api.SubjectApi.UpdateSubject)    // обновляет данные предмета
		subjectRoutes.DELETE("/:id", api.SubjectApi.DeleteSubject) // удаляет предмет
	}

	// Маршруты для учителей
	teacherRoutes := r.Group("/api/teacher")
	teacherRoutes.Use(middleware.AuthMiddleware(TeacherRoleID)) // middleware для аутентификации учителей
	{
		teacherRoutes.GET("/groups", api.ElseApi.GetCuratorGroupsStudentList)       // возвращает список групп, где учителю нужно поставить оценки
		teacherRoutes.GET("/studentAttendance/:id", api.ElseApi.StudentsAttendance) // возвращает список студентов с оценками
		teacherRoutes.POST("/studentAttendance", api.ElseApi.UpdateAttendance)      // обновляет оценки студентов
	}

	// Маршруты для администраторов
	adminRoutes := r.Group("/api/admin")
	adminRoutes.Use(middleware.AuthMiddleware(AdminRoleID)) // middleware для аутентификации администраторов
	{
		adminRoutes.GET("/AdminPanel", api.ElseApi.GetAdminPanelData) // возвращает данные для администратора
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

	//// Создаем пул соединений для PostgreSQL
	//dbPool, err := pgxpool.Connect(context.Background(), connToDb)
	//if err != nil {
	//	log.Fatal("Unable to connect to database:", err)
	//}
	//defer dbPool.Close()

	dbpu := users.NewUserPostgresHandlerDB(connToDb)
	dbru := users.NewUserHandlerRedis(os.Getenv("CONN_TO_REDIS"))
	apiUsers := handlers.NewUserHandler(dbpu, dbru)

	dbpr := role.NewRolePostgresHandler(connToDb)
	apiRoles := handlers.NewRoleHandler(dbpr)

	dbpg := groups.NewGroupPostgresRepository(connToDb)
	apiGroups := handlers.NewGroupHandler(dbpg)

	dbpn := note.NewNotePostgresHandlerDB(connToDb)
	apiNotes := handlers.NewNoteHandler(dbpn)

	dbps := schedules.NewSchedulePostgresHandlerDB(connToDb)
	apiSchedules := handlers.NewScheduleHandler(dbps)

	dbpsu := subjects.NewSubjectPostgresHandlerDB(connToDb)
	apiSubjects := handlers.NewSubjectHandler(dbpsu)

	dbpa := attendance.NewAttendancePostgresHandlerDB(connToDb)

	apiElse := handlers.NewElseHandler(dbpu, dbpg, dbpr, dbpa)

	api := ApiHandlers{UserApi: apiUsers, RoleApi: apiRoles,
		GroupApi: apiGroups, NoteApi: apiNotes,
		ScheduleApi: apiSchedules, SubjectApi: apiSubjects,
		ElseApi: apiElse}

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

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
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
