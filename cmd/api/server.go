package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/F1sssss/go_ecom/services/user"
	"github.com/gorilla/mux"
)

type Server struct {
	address string
	db      *sql.DB
}

func NewServer() *Server {
	return &Server{
		address: ":8080",
		db:      nil,
	}
}

func (s *Server) Start() error {

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	log.Println("Starting server on port", s.address)
	return http.ListenAndServe(s.address, nil)
}
