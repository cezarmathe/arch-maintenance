package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/cezarmathe/arch-maintenance/communication"
	"github.com/cezarmathe/arch-maintenance/config"
	"github.com/cezarmathe/arch-maintenance/maintenance"
)

var (

	log *logrus.Entry

	com chan uint
)

func main() {

	if os.Getenv("DEBUG") == "true" {
		log = (&logrus.Logger{
			Formatter:    &logrus.TextFormatter{FullTimestamp: true},
			Level: logrus.DebugLevel ,
			Out:          os.Stdout,
			ReportCaller: false,
		}).WithField("Component", "MAIN")
	} else {
		log = (&logrus.Logger{
			Formatter:    &logrus.TextFormatter{FullTimestamp: true},
			Level: logrus.InfoLevel ,
			Out:          os.Stdout,
			ReportCaller: false,
		}).WithField("Component", "MAIN")
	}

	config.LoadConfig()

	log.Debug("starting maintenance")
	go maintenance.Maintenance()

	log.Debug("creating the communication channel")
	com = make(chan uint)
	log.Debug("created the communication channel")
	maintenance.SetChannel(com)
	com <- communication.ChannelSet

	var msg uint = <- com

	if (msg != communication.ChannelReceive) {
		log.Panic(msg)
	}


}
