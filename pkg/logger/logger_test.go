package logger

import (
	"testing"
)

func TestJSONLogger(t *testing.T) {
	InitZapLogger("/Users/anthonyzero/logs/go-quick-api", ToLevel("warn"))
	Info("info message")
	Debug("debug message")
	Warnf("%s message", "warn")
	Error("error mssage")
	Panic("panic message")
}
