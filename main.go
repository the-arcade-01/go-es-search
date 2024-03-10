package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router *chi.Mux
}

func CreateServer() *Server {
	return &Server{
		Router: chi.NewRouter(),
	}
}

func (server *Server) MountMiddlewares() {
	server.Router.Use(middleware.Logger)
}

func (server *Server) MountRoutes() {
	server.Router.Get("/", greet)
	server.Router.Get("/{name}", greetName)
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Hello, World!!")
}

func greetName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmt.Sprintf("Hello, %v!!", name))
}

func main() {
	server := CreateServer()
	server.MountMiddlewares()
	server.MountRoutes()
	log.Println("server running on port:5000")
	http.ListenAndServe(":5000", server.Router)
}
