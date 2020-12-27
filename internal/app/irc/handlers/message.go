package handlers

import (
	"fmt"
	"strings"

	irc "github.com/fluffle/goirc/client"
	"github.com/hbjydev/cloudbot/internal/commands"
)

func Message(conn *irc.Conn, line *irc.Line) {
	nick := line.Nick
	content := line.Text()

	if strings.HasPrefix(content, "!") {
		args := strings.Split(content, " ")
		commandName := strings.Replace(args[0], "!", "", 1)

		switch commandName {
		case "github":
			fmt.Println("Handling command `github`...")
			conn.Privmsg(
				line.Target(),
				commands.GithubCommand(nick, line.Target(), args[1:]),
			)
		}
	}
}
