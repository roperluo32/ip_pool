package logrusger

import (
	"testing"
)

func TestLogrusLoggerBasic(t *testing.T) {
	logrusGer := NewLogrusLogger()
	logrusGer.Debug("hello world")
}
