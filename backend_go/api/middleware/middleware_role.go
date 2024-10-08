package middleware

import (
	"go.uber.org/zap"
)

type roleMiddleware struct {
	logger *zap.Logger
}

//func (r *roleMiddleware) MiddlewareZap() interface{} {
//	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
//		LogURI:    true,
//		LogStatus: true,
//		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
//			r.logger.Info("request",
//				zap.String("URI", v.URI),
//				zap.Int("status", v.Status),
//			)
//
//			return nil
//		},
//	})
//}
