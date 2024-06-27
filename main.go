package main

import (
	"github.com/the-arcade-01/go-es-search/server"
)

func main() {
	server := server.NewServer()
	server.Start()
}
