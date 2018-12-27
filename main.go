package main

import (
	"github.com/cezarmathe/arch-maintenance/config"
	"github.com/cezarmathe/arch-maintenance/logging"
)

func main() {
	config.SetLogger(logging.Loggers["config"])
	config.LoadConfig()


}
