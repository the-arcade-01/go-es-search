package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router *chi.Mux
}

func NewServer() *Server {
	server := &Server{
		Router: chi.NewRouter(),
	}
	server.MountMiddlewares()
	server.MountHandlers()
	return server
}

func (s *Server) MountMiddlewares() {
	s.Router.Use(middleware.Logger)
}

func (s *Server) MountHandlers() {
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
}

func (s *Server) Start() {
	log.Println("Server is running on :8080")
	http.ListenAndServe(":8080", s.Router)
}
