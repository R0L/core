package logger_test

import (
	"github.com/R0L/core/library/logger"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestInfo(t *testing.T) {
	logger.Info("1234567890")
}
