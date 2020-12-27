package app

import (
	"crypto/tls"
	"fmt"

	"github.com/hbjydev/cloudbot/internal/app/irc/handlers"
	"github.com/hbjydev/cloudbot/internal/pkg/config"

	irc "github.com/fluffle/goirc/client"
)

func CreateConnection() {
	// Create the IRC connection
	cfg := irc.NewConfig(config.IrcNick)
	cfg.SSL = true
	cfg.SSLConfig = &tls.Config{ServerName: "irc.freenode.net"}
	cfg.Server = "irc.freenode.net:7000" // TODO: Make this configurable
	cfg.NewNick = func(n string) string { return n + "^" }

	// Generate an IRC client
	c := irc.Client(cfg)

	// Set up IRC handlers
	c.HandleFunc(irc.CONNECTED, handlers.Connected)
	c.HandleFunc(irc.DISCONNECTED, handlers.Disconnected)
	c.HandleFunc(irc.PRIVMSG, handlers.Message)

	// Connect to the IRC server
	if err := c.Connect(); err != nil {
		fmt.Printf("Connection error: %s\n", err.Error())
	}
}
