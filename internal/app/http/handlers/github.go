package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hbjydev/cloudbot/internal/app/vars"
	"github.com/hbjydev/cloudbot/internal/pkg/config"
)

type webhook struct {
  Repository struct {
    ID       string `json:"id"`
    FullName string `json:"full_name"`
  } `json:"repository"`
}

type pushWebhook struct {
  webhook

  Ref string
  Before string
  After string
}

func GithubHandler(w http.ResponseWriter, r *http.Request) {
  dataBytes, err := ioutil.ReadAll(r.Body)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprint(w, "Something went wrong!")
    return
  }

  hookType := r.Header.Get("X-GitHub-Event")

  switch hookType {
    case "push":
      var hookData pushWebhook

      json.Unmarshal(dataBytes, &hookData)

      vars.Connection.Privmsg(
        config.IrcChannel,
        fmt.Sprintf(
          "New push to %s (%s): %s -> %s",
          hookData.Repository.FullName,
          hookData.Ref,
          hookData.Before[0:6],
          hookData.After[0:6],
        ),
      )
  }

  w.WriteHeader(http.StatusOK)
  fmt.Fprint(w, "Hello, world!")
}
