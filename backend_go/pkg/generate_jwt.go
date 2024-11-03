package pkg

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Secret key for signing the JWT (replace with env variable in production)
var jwtSecret = []byte("your_secret_key")

// GenerateJWT генерирует JWT-токен для пользователя.
func GenerateJWT(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // 72 часа жизни токена

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseJWT проверяет JWT-токен и возвращает userID, если токен действителен.
func ParseJWT(tokenStr string) (int, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, jwt.ErrSignatureInvalid
	}

	userID, err := strconv.Atoi(claims["user_id"].(string))
	if err != nil {
		return 0, err
	}

	return userID, nil
}
