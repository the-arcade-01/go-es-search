package main

import (
	"log"
	"net/http"

	"github.com/the-arcade-01/go-es-search/server/internal/api"
)

func main() {
	server := api.CreateServer()

	log.Println("[Main] server started running on port:8080")
	http.ListenAndServe(":8080", server.Router)
}
