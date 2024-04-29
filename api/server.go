package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hasssanezzz/rest-workers/storage"
)

type Server struct {
	listenAddr string
	storage    *storage.Storage
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		storage:    storage.NewStorage(),
	}
}

func (s *Server) Start() error {
	router := mux.NewRouter()

	router.HandleFunc("/", s.handleHelloWorld).Methods("GET")

	return http.ListenAndServe(s.listenAddr, router)
}

func (s *Server) handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, 200, "Hello world")
}
