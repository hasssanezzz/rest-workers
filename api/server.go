package api

import (
	"encoding/json"
	"net/http"
	"strconv"

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

	router.HandleFunc("/api/v0/task", s.handleListTasks).Methods("GET")
	router.HandleFunc("/api/v0/task/{id}", s.handleGetTask).Methods("GET")
	router.HandleFunc("/api/v0/task", s.handleCreateTask).Methods("POST")
	router.HandleFunc("/api/v0/task/{id}", s.handleDeleteTask).Methods("DELETE")

	return http.ListenAndServe(s.listenAddr, router)
}

func (s *Server) handleListTasks(w http.ResponseWriter, r *http.Request) {
	tasks := s.storage.List()
	WriteJSON(w, 200, tasks)
}

func (s *Server) handleGetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		// TODO handle error: invalid id
		WriteJSON(w, 400, "bad ID")
		return
	}

	tasks, err := s.storage.Get(taskId)
	if err != nil {
		// TODO handle error: 404
		WriteJSON(w, 404, "task not found")
		return
	}

	WriteJSON(w, 200, tasks)
}

func (s *Server) handleCreateTask(w http.ResponseWriter, r *http.Request) {}

func (s *Server) handleDeleteTask(w http.ResponseWriter, r *http.Request) {}
