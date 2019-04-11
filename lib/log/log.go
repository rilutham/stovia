package log

import (
	"fmt"
	"io"
	"path/filepath"
	"runtime"

	log "github.com/sirupsen/logrus"
)

var std = log.New()

// Level is an alias type to the logger implementation
type Level log.Level

// Format :nodoc:
type Format uint8

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	OFF
)

const (
	_ Format = iota
	JSONFormat
	TextFormat
)

// Logger is an interface for general logging
type Logger interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Debug(...interface{})
	Debugf(string, ...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
	Warn(...interface{})
	Warnf(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})
}

// Context is a function to store detail code line error
type Context struct {
	Package string
	Scope   string
	Line    int
	File    string
}

// GetLogger is a function to get default logger
func GetLogger() Logger {
	return std
}

// SetOutput to change logger argsput
func SetOutput(w io.Writer) {
	std.Out = w
}

// SetLevel of the logger
func SetLevel(level Level) {
	std.Level = log.Level(level)
}

// SetFormat for the logger
func SetFormat(format Format) {
	switch format {
	case JSONFormat:
		std.Formatter = &log.JSONFormatter{}
	default:
		std.Formatter = &log.TextFormatter{}
	}
}

// Print is an alias method to the logger implementation
func Print(args ...interface{}) {
	std.Print(args...)
}

// Printf is an alias method to the logger implementation
func Printf(format string, args ...interface{}) {
	std.Printf(format, args...)
}

// Debug is an alias method to the logger implementation
func Debug(args ...interface{}) {
	std.Debug(args...)
}

// Debugf is an alias method to the logger implementation
func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

// Info is an alias method to the logger implementation
func Info(args ...interface{}) {
	std.Info(args...)
}

// Infof is an alias method to the logger implementation
func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

// Warn is an alias method to the logger implementation
func Warn(args ...interface{}) {
	std.Warn(args...)
}

// Warnf is an alias method to the logger implementation
func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}

// Error is an alias method to the logger implementation
func Error(args ...interface{}) {
	std.Error(args...)
}

// Errorf is an alias method to the logger implementation
func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

// Fatal is an alias method to the logger implementation
func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

// Fatalf is an alias method to the logger implementation
func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

// Standard :nodoc:
func Standard() *log.Logger {
	return std
}

// For :nodoc:
func For(pkg, scope string) Logger {
	_, file, line, _ := runtime.Caller(1)
	_, file = filepath.Split(file)
	return mWithContext(Context{
		Package: pkg,
		Scope:   scope,
		Line:    line,
		File:    file,
	})
}

func mWithContext(c Context) *log.Entry {
	field := log.Fields{
		"package": c.Package,
		"scope":   fmt.Sprintf("%s[%s:%d]", c.Scope, c.File, c.Line),
	}

	return std.WithFields(field)
}
