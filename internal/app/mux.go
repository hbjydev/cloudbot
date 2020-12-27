package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hbjydev/cloudbot/internal/app/http/handlers"
)

func CreateMuxServer() {
  server := mux.NewRouter()

  server.HandleFunc("/github", handlers.GithubHandler)

  log.Fatal(http.ListenAndServe(":3000", server))
}
