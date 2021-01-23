package log

import (
	"log"
	"os"
)

// Level is an int enum with log levels,
// "none", "err", "info" (default), "debug" are initialised as iota
type Level int

const (
	none Level = iota
	err
	info
	debug
)

var levels = map[string]Level{
	"NONE":  none,
	"ERR":   err,
	"INFO":  info,
	"DEBUG": debug,
}

// LoggerOptions specifies options that can be used in logger
type LoggerOptions struct {
	level Level
}

var (
	errorLogger = log.New(os.Stderr, "ERROR: ", log.LstdFlags)
	infoLogger  = log.New(os.Stdout, "INFO: ", log.LstdFlags)
	debugLogger = log.New(os.Stdout, "DEBUG: ", log.LstdFlags)
)

// Err shows error while processing application
func Err(args ...interface{}) {
	if options.level < err {
		return
	}

	errorLogger.Fatalln(args...)
}

// Info shows main information about processing application data
func Info(args ...interface{}) {
	if options.level < info {
		return
	}

	infoLogger.Println(args...)
}

// Debug shows more detailed information about processing application data
func Debug(args ...interface{}) {
	if options.level < debug {
		return
	}

	debugLogger.Println(args...)
}

func readLevel(inputLevel string) Level {
	if level, ok := levels[inputLevel]; ok {
		return level
	}

	return info
}

// Configure allows to configure application logger
// When initialised, it reads `LOG_LEVEL` variable
// Use `LOG_LEVEL` environment variable - "none", "err", "info" (default), "debug"
func Configure(inputLevel string) LoggerOptions {
	var (
		level   = readLevel(inputLevel)
		options = LoggerOptions{level}
	)

	return options
}

var options = Configure(os.Getenv("LOG_LEVEL"))
