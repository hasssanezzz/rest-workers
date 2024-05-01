package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func (s *Server) Start() error {
	router := mux.NewRouter()

	router.HandleFunc("/api/v0/task", s.handleListTasks).Methods("GET")
	router.HandleFunc("/api/v0/task/{id}", s.handleGetTask).Methods("GET")
	router.HandleFunc("/api/v0/task", s.handleCreateTask).Methods("POST")

	return http.ListenAndServe(s.listenAddr, router)
}
