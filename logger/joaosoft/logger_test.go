package logger_joaosoft

import (
	"testing"

	logger "github.com/joaosoft/logger"
)

func BenchmarkJoaosoftLoggerWithFormatedMessage(b *testing.B) {
	var log = logger.NewLogDefault("joaosoft-logger", logger.InfoLevel).WithTag("user", "account")

	for i := 0; i < b.N; i++ {
		log.Infof("hello %d", i)
	}
}

func BenchmarkJoaosoftLoggerWithSimpleMessage(b *testing.B) {
	var log = logger.NewLogDefault("joaosoft-logger", logger.InfoLevel).WithTag("user", "account")

	for i := 0; i < b.N; i++ {
		log.Info("hello")
	}
}
