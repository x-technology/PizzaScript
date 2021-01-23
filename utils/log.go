package utils

import (
	"log"
	"os"
)

type LogLevel int

const (
	NONE LogLevel = iota
	ERR
	INFO
	DEBUG
)

var levels = map[string]LogLevel{
	"NONE":  NONE,
	"ERR":   ERR,
	"INFO":  INFO,
	"DEBUG": DEBUG,
}

// loggerOptions
type loggerOptions struct {
	level LogLevel
}

type logger interface {
	Err()
	Debug()
	Info()
}

// Logger is created as singletone for application
type Logger struct {
	options loggerOptions
}

var (
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Lshortfile)
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Lshortfile)
	debugLogger = log.New(os.Stdout, "DEBUG: ", log.Lshortfile)
)

// Err shows error while processing application
func (l *Logger) Err(args ...interface{}) {
	if l.options.level < ERR {
		return
	}

	errorLogger.Fatalln(args...)
}

// Info shows main information about processing application data
func (l *Logger) Info(args ...interface{}) {
	if l.options.level < INFO {
		return
	}

	infoLogger.Println(args...)
}

// Debug shows more detailed information about processing application data
func (l *Logger) Debug(args ...interface{}) {
	if l.options.level < DEBUG {
		return
	}

	debugLogger.Println(args...)
}

func readLevel(inputLevel string) LogLevel {
	if level, ok := levels[inputLevel]; ok {
		return level
	}

	return INFO
}

// Use `LOG_LEVEL` environment variable - "none", "err", "info" (default), "debug"
var (
	level   = readLevel(os.Getenv("LOG_LEVEL"))
	options = loggerOptions{level}
)

// Log to simplify logging levels on application
var Log = &Logger{options}
