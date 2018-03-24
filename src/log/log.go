package log

import (
	logging "github.com/op/go-logging" // A logging library
)

var (
	log *logging.Logger
)

// -----------------------
// Initialization function
// -----------------------

// Init starts the logger
func Init() {
	log = logging.MustGetLogger("nd")
}

// Info prints to the console
func Info(args ...interface{}) {
	log.Info(args...)
}

// Fatal ends the application
func Fatal(message string, err error) {
	log.Fatal(message, err)
}

// Error does more than print to the console
func Error(args ...interface{}) {
	log.Error(args...)
}
