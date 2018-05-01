package gobenchmark

import (
	"testing"

	logger "gitlab.com/vredens/go-logger"
)

func BenchmarkLoggerWithFormatedMessage(b *testing.B) {
	var log = logger.Spawn(
		logger.WithFields(map[string]interface{}{"component": "logger"}),
		logger.WithTags("user", "account"),
	)
	logger.Reconfigure(logger.WithFormat(logger.FormatJSON))

	for i := 0; i < b.N; i++ {
		log.Infof("hello %d", i)
	}
}

func BenchmarkLoggerWithSimpleMessage(b *testing.B) {
	var log = logger.Spawn(
		logger.WithFields(map[string]interface{}{"component": "logger"}),
		logger.WithTags("user", "account"),
	)
	logger.Reconfigure(logger.WithFormat(logger.FormatJSON))

	for i := 0; i < b.N; i++ {
		log.Info("hello")
	}
}
