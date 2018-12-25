package logging

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var (
// logger = log.New()
)

func init() {
	log.SetReportCaller(true)

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	log.SetOutput(os.Stdout)
}

type Logger interface {
	Trace()
	Debug()
	Info()
	Warn()
	Error()
	Fatal()
	Panic()
}
