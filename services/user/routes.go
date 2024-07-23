package user

import (
	"net/http"

	"github.com/F1sssss/go_ecom/types"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", h.getUserByID).Methods("GET")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// Create a new user
}

func (h *Handler) getUserByID(w http.ResponseWriter, r *http.Request) {
	// Get a user by ID
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// Get a user by username
}
