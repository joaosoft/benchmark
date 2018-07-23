package logger_vredens

import (
	"testing"

	logger "gitlab.com/vredens/go-logger"
)

func BenchmarkVredensLoggerWithFormatedMessage(b *testing.B) {
	var log = logger.Spawn(
		logger.WithFields(map[string]interface{}{"component": "vredens-logger"}),
		logger.WithTags("user", "account"),
	)
	logger.Reconfigure(logger.WithFormat(logger.FormatJSON))

	for i := 0; i < b.N; i++ {
		log.Infof("hello %d", i)
	}
}

func BenchmarkVredensLoggerWithSimpleMessage(b *testing.B) {
	var log = logger.Spawn(
		logger.WithFields(map[string]interface{}{"component": "logger"}),
		logger.WithTags("user", "account"),
	)
	logger.Reconfigure(logger.WithFormat(logger.FormatJSON))

	for i := 0; i < b.N; i++ {
		log.Info("hello")
	}
}
