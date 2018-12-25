package logging

import "github.com/sirupsen/logrus"

type configLogger struct {
	log *logrus.Entry
}

func newConfigLogger() *configLogger {
	return &configLogger{
		log: logger.WithFields(logrus.Fields{
			"Component": "CONFIG",
		}),
	}
}

func (c *configLogger) Exist() {}

func (c *configLogger) ConfigFileLoadStatus(filepath string, status bool) {
	log.WithFields(logrus.Fields{
		"filepath": filepath,
		"status":   status,
	}).Info("Load config from file")
}
