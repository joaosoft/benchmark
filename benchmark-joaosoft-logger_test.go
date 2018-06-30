package benchmark

import (
	"testing"

	golog "github.com/joaosoft/go-log/app"
)

func BenchmarkGoLogWithFormatedMessage(b *testing.B) {
	var log = golog.NewLogDefault("golog", golog.InfoLevel).WithTag("user", "account")

	for i := 0; i < b.N; i++ {
		log.Infof("hello %d", i)
	}
}

func BenchmarkGoLogWithSimpleMessage(b *testing.B) {
	var log = golog.NewLogDefault("golog", golog.InfoLevel).WithTag("user", "account")

	for i := 0; i < b.N; i++ {
		log.Info("hello")
	}
}
