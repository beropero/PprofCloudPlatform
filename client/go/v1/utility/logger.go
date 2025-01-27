package utils

import (
    "fmt"
    "log"
)

type Logger interface {
    Debug(msg string, args ...interface{})
    Info(msg string, args ...interface{})
    Warn(msg string, args ...interface{})
    Error(msg string, args ...interface{})
}

type DefaultLogger struct {
    debug bool
}

func NewDefaultLogger(debug bool) *DefaultLogger {
    return &DefaultLogger{debug: debug}
}

func (l *DefaultLogger) log(level, msg string, args ...interface{}) {
    entry := fmt.Sprintf("[%s] %s", level, msg)
    if len(args) > 0 {
        entry = fmt.Sprintf(entry, args...)
    }
    log.Println(entry)
}

func (l *DefaultLogger) Debug(msg string, args ...interface{}) {
    if l.debug {
        l.log("DEBUG", msg, args...)
    }
}

func (l *DefaultLogger) Info(msg string, args ...interface{}) {
    l.log("INFO", msg, args...)
}

func (l *DefaultLogger) Warn(msg string, args ...interface{}) {
    l.log("WARN", msg, args...)
}

func (l *DefaultLogger) Error(msg string, args ...interface{}) {
    l.log("ERROR", msg, args...)
}