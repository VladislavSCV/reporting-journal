package middleware

import (
	"go.uber.org/zap"
)

type authMiddleware struct {
	logger *zap.Logger
}
