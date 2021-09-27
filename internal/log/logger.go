package log

import (
	"io"
	"log"
	"os"
)

// Logger 日志结构体
type Logger struct {
	*log.Logger
	level Level
}

// New 新的日志
func New(out io.Writer, prefix string, flag int, level Level) *Logger {
	return &Logger{
		Logger: log.New(out, prefix, flag),
		level:  level,
	}
}

var (
	std  = New(os.Stdout, DebugPrefix, log.LstdFlags, LogDebug)
	info = New(os.Stdout, InfoPrefix, log.LstdFlags, LogInfo)
	warn = New(os.Stdout, WarnPrefix, log.LstdFlags, LogWarn)
)

// Level 日志级别
type Level int

const (
	// LogDebug 调试模式
	LogDebug Level = iota

	// LogInfo 日志信息
	LogInfo

	// LogWarn 警告
	LogWarn
)

const (
	// InfoPrefix 信息前缀
	InfoPrefix = "[I] "

	// DebugPrefix 调试前缀
	DebugPrefix = "[D] "

	// WarnPrefix 警告前缀
	WarnPrefix = "[W] "
)

// Info 信息
func (l *Logger) Info(v ...interface{}) {
	l.output(LogInfo, v...)
}

func (l *Logger) output(level Level, v ...interface{}) {
	if l.level <= level {
		l.Logger.Println(v...)
	}
}

// Debug 调试
func (l *Logger) Debug(v ...interface{}) {
	l.output(LogDebug, v...)
}

// Warn 警告
func (l *Logger) Warn(v ...interface{}) {
	l.output(LogWarn, v...)
}

// Debug 调试
func Debug(v ...interface{}) {
	std.Debug(v...)
}

// Info 信息
func Info(v ...interface{}) {
	info.Info(v...)
}

// Warn 警告
func Warn(v ...interface{}) {
	warn.Warn(v...)
}
