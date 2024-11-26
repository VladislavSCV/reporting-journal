package middleware

import (
	"github.com/VladislavSCV/pkg"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware проверяет токен и роль пользователя
func AuthMiddleware(requiredRoleID int) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Извлечение токена из заголовка Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			//logger.Error("missing Authorization header")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}

		// Проверка формата заголовка, например: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			//logger.Error("invalid Authorization header format")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}
		token := parts[1]

		// Парсинг JWT токена
		userId, userRoleID, err := pkg.ParseJWT(token)
		if err != nil {
			//logger.Error("failed to parse JWT", zap.Error(err))
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Проверка на наличие userRoleID
		if userRoleID == 0 {
			//logger.Error("empty user role ID from JWT")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Проверка роли пользователя
		if userRoleID != requiredRoleID && userRoleID != 3 {
			//logger.Error("user does not have required role",
			//	zap.Int("required_role_id", requiredRoleID),
			//	zap.Int("user_role_id", userRoleID),
			//)
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		// Сохранение в контекст
		c.Set("user_id", userId)

		c.Set("user_role_id", userRoleID)

		// Передача управления следующему обработчику
		c.Next()
	}
}
