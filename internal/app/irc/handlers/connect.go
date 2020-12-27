package handlers

import (
	"fmt"

	irc "github.com/fluffle/goirc/client"
	"github.com/hbjydev/cloudbot/internal/app/vars"
)

func Connected(conn *irc.Conn, line *irc.Line) {
	fmt.Println("Connected to IRC server")
  vars.Connection = conn
	conn.Join("#cloudbot-devel")
}
