package app

import (
	"fmt"

	"github.com/hbjydev/cloudbot/internal/app/vars"
)

func Startup() {
	fmt.Println("Cloudbot has reached startup...")
	CreateConnection()
  CreateMuxServer()

	// Wait to be told to quit
	<-vars.Quit
}
