package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	// the logger
	logger *logrus.Logger

	// Loggers is a map that contains loggers for each component
	// Loggers map[string]*logrus.Entry
	Loggers map[string]Logger
)

func init() {
	Loggers = make(map[string]Logger)

	logger = &logrus.Logger{
		Formatter:    &logrus.TextFormatter{FullTimestamp: true},
		Level:        logrus.TraceLevel,
		Out:          os.Stdout,
		ReportCaller: false,
	}

	// populating the loggers for each component
	Loggers["config"] = newConfigLogger()
}

// Logger defines the functions that every logger should implement
type Logger interface {
	Exist()
}
