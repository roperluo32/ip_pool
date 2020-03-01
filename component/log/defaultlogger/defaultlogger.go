package defaultlogger

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
}

// DefaultLogger 默认的logger，使用系统的log实现
type DefaultLogger struct {
}

// Trace logs a message at level Trace on the standard logger.
func (dl *DefaultLogger) Trace(args ...interface{}) {
	log.SetPrefix("[TRACE] ")
	log.Output(3, fmt.Sprintln(args...))
}

// Debug logs a message at level Debug on the standard logger.
func (dl *DefaultLogger) Debug(args ...interface{}) {
	log.SetPrefix("[DEBUG] ")
	log.Output(3, fmt.Sprintln(args...))
}

// Print logs a message at level Info on the standard logger.
func (dl *DefaultLogger) Print(args ...interface{}) {
	log.SetPrefix("")
	log.Output(3, fmt.Sprintln(args...))
}

// Info logs a message at level Info on the standard logger.
func (dl *DefaultLogger) Info(args ...interface{}) {
	log.SetPrefix("[INFO] ")
	log.Output(3, fmt.Sprintln(args...))
}

// Warn logs a message at level Warn on the standard logger.
func (dl *DefaultLogger) Warn(args ...interface{}) {
	log.SetPrefix("[WARN] ")
	log.Output(3, fmt.Sprintln(args...))
}

// Error logs a message at level Error on the standard logger.
func (dl *DefaultLogger) Error(args ...interface{}) {
	log.SetPrefix("[ERROR] ")
	log.Output(3, fmt.Sprintln(args...))
}

// Panic logs a message at level Panic on the standard logger.
func (dl *DefaultLogger) Panic(args ...interface{}) {
	log.SetPrefix("[PANIC] ")
	log.Output(3, fmt.Sprintln(args...))
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func (dl *DefaultLogger) Fatal(args ...interface{}) {
	log.SetPrefix("[FATAL] ")
	log.Output(3, fmt.Sprintln(args...))
}

// Tracef logs a message at level Trace on the standard logger.
func (dl *DefaultLogger) Tracef(format string, args ...interface{}) {
	log.SetPrefix("[TRACE] ")
	log.Output(3, fmt.Sprintf(format, args...))
}

// Debugf logs a message at level Debug on the standard logger.
func (dl *DefaultLogger) Debugf(format string, args ...interface{}) {
	log.SetPrefix("[DEBUG] ")
	log.Output(3, fmt.Sprintf(format, args...))
}

// Printf logs a message at level Info on the standard logger.
func (dl *DefaultLogger) Printf(format string, args ...interface{}) {
	log.SetPrefix("")
	log.Output(3, fmt.Sprintf(format, args...))
}

// Infof logs a message at level Info on the standard logger.
func (dl *DefaultLogger) Infof(format string, args ...interface{}) {
	log.SetPrefix("[INFO] ")
	log.Output(3, fmt.Sprintf(format, args...))
}

// Warnf logs a message at level Warn on the standard logger.
func (dl *DefaultLogger) Warnf(format string, args ...interface{}) {
	log.SetPrefix("[WARN] ")
	log.Output(3, fmt.Sprintf(format, args...))
}

// Errorf logs a message at level Error on the standard logger.
func (dl *DefaultLogger) Errorf(format string, args ...interface{}) {
	log.SetPrefix("[ERROR] ")
	log.Output(3, fmt.Sprintf(format, args...))
}

// Panicf logs a message at level Panic on the standard logger.
func (dl *DefaultLogger) Panicf(format string, args ...interface{}) {
	log.SetPrefix("[PANIC] ")
	log.Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func (dl *DefaultLogger) Fatalf(format string, args ...interface{}) {
	log.SetPrefix("[FATAL] ")
	log.Fatalf(format, args...)
}
