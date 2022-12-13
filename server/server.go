package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Chahine-tech/ws/db/db"
	"github.com/Chahine-tech/ws/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

type Server struct {
	Router  *chi.Mux
	Queries *db.Queries
}

func NewServer() *Server {

	s := &Server{}

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cant log config")
	}

	dbSource := config.DBSource
	dbDriver := config.DBDriver
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to the db : ", err.Error())

	}
	s.Queries = db.New(conn)
	log.Println("Connected to the database")

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
