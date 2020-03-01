package logrusger

import (
	"ip_proxy/component/config"
	"sync"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	caller "github.com/xdxiaodong/logrus-hook-caller"
)

type LogrusGer struct {
	inst *logrus.Logger
}

// Trace logs a message at level Trace on the standard logger.
func (lg *LogrusGer) Trace(args ...interface{}) {
	lg.inst.Trace(args...)
}

// Debug logs a message at level Debug on the standard logger.
func (lg *LogrusGer) Debug(args ...interface{}) {
	lg.inst.Debug(args...)
}

// Print logs a message at level Info on the standard logger.
func (lg *LogrusGer) Print(args ...interface{}) {
	lg.inst.Print(args...)
}

// Info logs a message at level Info on the standard logger.
func (lg *LogrusGer) Info(args ...interface{}) {
	lg.inst.Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func (lg *LogrusGer) Warn(args ...interface{}) {
	lg.inst.Warn(args...)
}

// Error logs a message at level Error on the standard logger.
func (lg *LogrusGer) Error(args ...interface{}) {
	lg.inst.Error(args...)
}

// Panic logs a message at level Panic on the standard logger.
func (lg *LogrusGer) Panic(args ...interface{}) {
	lg.inst.Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func (lg *LogrusGer) Fatal(args ...interface{}) {
	lg.inst.Fatal(args...)
}

// Tracef logs a message at level Trace on the standard logger.
func (lg *LogrusGer) Tracef(format string, args ...interface{}) {
	lg.inst.Tracef(format, args...)
}

// Debugf logs a message at level Debug on the standard logger.
func (lg *LogrusGer) Debugf(format string, args ...interface{}) {
	lg.inst.Debugf(format, args...)
}

// Printf logs a message at level Info on the standard logger.
func (lg *LogrusGer) Printf(format string, args ...interface{}) {
	lg.inst.Printf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func (lg *LogrusGer) Infof(format string, args ...interface{}) {
	lg.inst.Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func (lg *LogrusGer) Warnf(format string, args ...interface{}) {
	lg.inst.Warnf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func (lg *LogrusGer) Errorf(format string, args ...interface{}) {
	lg.inst.Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func (lg *LogrusGer) Panicf(format string, args ...interface{}) {
	lg.inst.Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func (lg *LogrusGer) Fatalf(format string, args ...interface{}) {
	lg.inst.Fatalf(format, args...)
}

var once sync.Once
var _logrusLogger LogrusGer

func initFormatter(ins *logrus.Logger) {
	formatter := new(prefixed.TextFormatter)
	formatter.FullTimestamp = true
	formatter.SetColorScheme(&prefixed.ColorScheme{
		PrefixStyle:    "blue+b",
		TimestampStyle: "white+h",
	})
	ins.SetFormatter(formatter)
}

// NewLogrusLogger 新建一个LogrusGer实例
func NewLogrusLogger() *LogrusGer {
	once.Do(func() {
		ins := logrus.New()
		var levleMap = map[string]logrus.Level{
			"trace": logrus.TraceLevel,
			"debug": logrus.DebugLevel,
			"info":  logrus.InfoLevel,
			"warn":  logrus.WarnLevel,
			"error": logrus.ErrorLevel,
			"fatal": logrus.FatalLevel,
			"panic": logrus.PanicLevel,
		}
		ins.SetLevel(levleMap[config.C.Log.Level])
		initFormatter(ins)
		hookCaller := caller.NewHook(&caller.CallerHookOptions{
			Flags: caller.Llongfile | caller.LstdFlags,
		})
		ins.AddHook(hookCaller)

		_logrusLogger = LogrusGer{
			inst: ins,
		}
	})

	return &_logrusLogger
}
