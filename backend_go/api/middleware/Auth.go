package middleware

import (
	"github.com/VladislavSCV/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RoleMiddleware проверяет наличие нужной роли
func RoleMiddleware(allowedRoleIDs []int) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Извлекаем токен из заголовка
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing"})
			return
		}

		// Проверяем токен и извлекаем данные
		claims, err := pkg.VerifyToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Проверяем роль пользователя
		for _, allowedRoleID := range allowedRoleIDs {
			if claims.RoleID == allowedRoleID {
				c.Set("userID", claims.UserID) // Сохраняем userID в контексте для дальнейшего использования
				c.Next()
				return
			}
		}

		// Если роль не подходит
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
	}
}

//r := gin.Default()
//
//adminRoutes := r.Group("/admin")
//adminRoutes.Use(middleware.RoleMiddleware([]int{1})) // Например, роль с ID = 1 — администратор
//{
//adminRoutes.GET("/dashboard", adminDashboardHandler)
//}
//
//userRoutes := r.Group("/user")
//userRoutes.Use(middleware.RoleMiddleware([]int{2})) // Например, роль с ID = 2 — пользователь
//{
//userRoutes.GET("/profile", userProfileHandler)
//}
//
//r.Run(":8080")
