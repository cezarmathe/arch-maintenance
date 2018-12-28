package maintenance

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/cezarmathe/arch-maintenance/communication"
)

var (
	com chan uint

	log *logrus.Entry
)

func init() {
	if os.Getenv("DEBUG") == "true" {
		log = (&logrus.Logger{
			Formatter:    &logrus.TextFormatter{FullTimestamp: true},
			Level: logrus.DebugLevel ,
			Out:          os.Stdout,
			ReportCaller: false,
		}).WithField("Component", "MAINTENANCE")
	} else {
		log = (&logrus.Logger{
			Formatter:    &logrus.TextFormatter{FullTimestamp: true},
			Level: logrus.InfoLevel ,
			Out:          os.Stdout,
			ReportCaller: false,
		}).WithField("Component", "MAINTENANCE")
	}
}

func SetChannel(_com chan uint) {
	com = _com
}

func Maintenance() {

	var msg uint

	log.Info("started maintenance")

	log.Debug("waiting to receive channel")
	msg = <-com

	if msg != communication.ChannelSet {
		log.Panic(msg)
	}

	log.Debug("received channel")
	com <- communication.ChannelReceive




}