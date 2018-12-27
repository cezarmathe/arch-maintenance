package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	// the logger
	logger *logrus.Logger

	// Loggers is a map that contains loggers for each component
	Loggers map[string]*logrus.Entry

	//the logging level
	loggingLevel logrus.Level
)

func init() {

	switch os.Getenv("ARCH_MAINTENANCE_LOGGING_LEVEL") {
	case "trace":
		loggingLevel = logrus.TraceLevel
	case "debug":
		loggingLevel = logrus.DebugLevel
	case "info":
		loggingLevel = logrus.InfoLevel
	case "warn":
		loggingLevel = logrus.WarnLevel
	case "fatal":
		loggingLevel = logrus.FatalLevel
	case "panic":
		loggingLevel = logrus.PanicLevel
	default:
		loggingLevel = logrus.TraceLevel
	}

	Loggers = make(map[string]*logrus.Entry)

	logger = &logrus.Logger{
		Formatter:    &logrus.TextFormatter{FullTimestamp: true},
		Level:        loggingLevel,
		Out:          os.Stdout,
		ReportCaller: false,
	}

	// populating the loggers for each component
	Loggers["config"] = logger.WithFields(logrus.Fields{
		"Component" : "CONFIG",
	})
}

