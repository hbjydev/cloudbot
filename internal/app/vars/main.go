package vars

import (
  irc "github.com/fluffle/goirc/client"
)

var (
	Quit = make(chan bool)

  Connection *irc.Conn
)
