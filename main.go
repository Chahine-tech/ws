package main

import (
	"github.com/Chahine-tech/ws/handlers"
	"github.com/Chahine-tech/ws/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	s := server.NewServer()
	r.Use(middleware.Logger)
	s.Router.Mount("/", r)
	s.Router.Get("/hello", handlers.HelloHandler())
	s.RunServer()
}
