package commands

import (
	"fmt"
	"strings"

	"github.com/hbjydev/cloudbot/internal/pkg/config"
)

func GithubCommand(nick string, channel string, args []string) string {
	repo := args[0]

  if !strings.Contains(repo, "/") {
    repo = config.GithubDefaultOwner + "/" + repo
  }

	return fmt.Sprintf("https://github.com/%s", repo)
}
