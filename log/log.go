package log

import (
	"context"
	"fmt"
	"log/slog"
	"runtime"
	"sync/atomic"
	"time"
)

const (
	LevelTrace slog.Level = slog.LevelDebug - 4
	LevelFatal slog.Level = slog.LevelError + 4
)

var defaultLogger atomic.Value

func init() {
	// Default logger has an extra skip to account for the log function
	defaultLogger.Store(&Slogger{
		logger: slog.Default(),
		skip:   4,
	})
}

func SetDefault(l Logger) {
	if sl, ok := l.(Slogger); ok {
		sl.skip = 4
	}
	defaultLogger.Store(l)
}

type Logger interface {
	Trace(args ...interface{})
	Tracef(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

type Slogger struct {
	logger *slog.Logger
	skip   int
}

func New() Logger {
	return &Slogger{
		logger: slog.Default(),
		skip:   3,
	}
}

func (l Slogger) log(level slog.Level, line string) {
	if !l.logger.Enabled(context.Background(), level) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(l.skip, pcs[:]) // Skip [Caller, log, XXXf]
	r := slog.NewRecord(time.Now(), level, line, pcs[0])
	_ = l.logger.Handler().Handle(context.Background(), r)
}

func (l Slogger) Trace(args ...interface{}) {
	l.log(slog.LevelDebug, fmt.Sprint(args...))
}
func (l Slogger) Tracef(format string, args ...interface{}) {
	l.log(slog.LevelDebug, fmt.Sprintf(format, args...))
}
func (l Slogger) Debug(args ...interface{}) {
	l.log(slog.LevelDebug, fmt.Sprint(args...))
}
func (l Slogger) Debugf(format string, args ...interface{}) {
	l.log(slog.LevelDebug, fmt.Sprintf(format, args...))
}
func (l Slogger) Info(args ...interface{}) {
	l.log(slog.LevelInfo, fmt.Sprint(args...))
}
func (l Slogger) Infof(format string, args ...interface{}) {
	l.log(slog.LevelInfo, fmt.Sprintf(format, args...))
}
func (l Slogger) Warn(args ...interface{}) {
	l.log(slog.LevelWarn, fmt.Sprint(args...))
}
func (l Slogger) Warnf(format string, args ...interface{}) {
	l.log(slog.LevelWarn, fmt.Sprintf(format, args...))
}
func (l Slogger) Error(args ...interface{}) {
	l.log(slog.LevelError, fmt.Sprint(args...))
}
func (l Slogger) Errorf(format string, args ...interface{}) {
	l.log(slog.LevelError, fmt.Sprintf(format, args...))
}
func (l Slogger) Fatal(args ...interface{}) {
	l.log(LevelFatal, fmt.Sprint(args...))
}
func (l Slogger) Fatalf(format string, args ...interface{}) {
	l.log(LevelFatal, fmt.Sprintf(format, args...))
}

func Trace(args ...interface{}) {
	Default().Trace(args...)
}
func Tracef(format string, args ...interface{}) {
	Default().Tracef(format, args...)
}
func Debug(args ...interface{}) {
	Default().Debug(args...)
}
func Debugf(format string, args ...interface{}) {
	Default().Debugf(format, args...)
}
func Info(args ...interface{}) {
	Default().Info(args...)
}
func Infof(format string, args ...interface{}) {
	Default().Infof(format, args...)
}
func Warn(args ...interface{}) {
	Default().Warn(args...)
}
func Warnf(format string, args ...interface{}) {
	Default().Warnf(format, args...)
}
func Error(args ...interface{}) {
	Default().Error(args...)
}
func Errorf(format string, args ...interface{}) {
	Default().Errorf(format, args...)
}
func Fatal(args ...interface{}) {
	Default().Fatal(args...)
}
func Fatalf(format string, args ...interface{}) {
	Default().Fatalf(format, args...)
}
