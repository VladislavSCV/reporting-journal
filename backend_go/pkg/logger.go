package pkg

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

// init initializes the global logger variable with a production logger.
// It panics if it is unable to create the logger.
func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

// Log logs the given data at the info level. It is a convenience wrapper
// around logger.Info and zap.Any.
func Log[T any](data T) {
	logger.Info("", zap.Any("data", data))
}
