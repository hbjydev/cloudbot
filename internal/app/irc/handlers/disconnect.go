package handlers

import (
	"fmt"

	"github.com/hbjydev/cloudbot/internal/app/vars"

	irc "github.com/fluffle/goirc/client"
)

func Disconnected(conn *irc.Conn, line *irc.Line) {
	fmt.Println("Disconnected from IRC server.")
	vars.Quit <- true
}
