package api

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/the-arcade-01/go-es-search/server/internal/service"
	"github.com/the-arcade-01/go-es-search/server/internal/utils"
)

type Server struct {
	Router *chi.Mux
}

func (server *Server) MountMiddlewares() {
	server.Router.Use(middleware.Heartbeat(utils.HEALTH))
	server.Router.Use(middleware.Logger)
}

func (server *Server) MountHandlers() {
	service := service.NewAPIService()
	server.Router.Get(utils.GREET, service.Greet)
}

func CreateServer() *Server {
	server := &Server{
		Router: chi.NewRouter(),
	}
	server.MountMiddlewares()
	server.MountHandlers()

	log.Println("[CreateServer] Server initialized")
	return server
}
