package main

import (
	// "github.com/cezarmathe/arch-maintenance/cmd"
	"fmt"
	"os"

	"github.com/cezarmathe/arch-maintenance/config"
	// log "github.com/cezarmathe/arch-maintenance/logging"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%+v", config.Config)

}
