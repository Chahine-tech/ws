package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router *chi.Mux
}

func NewServer() *Server {

	s := &Server{}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	s.Router = r
	return s
}

func (s *Server) RunServer() {

	fmt.Println("Serving on port 3000....")
	http.ListenAndServe(":3000", s.Router)
}
