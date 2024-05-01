package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/hasssanezzz/rest-workers/types"
)

func (s *Server) handleListTasks(w http.ResponseWriter, r *http.Request) {
	tasks := s.storage.List()
	WriteJSON(w, http.StatusOK, tasks)
}

func (s *Server) handleGetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		// TODO handle error: invalid id
		WriteJSON(w, http.StatusBadRequest, "bad ID")
		return
	}

	tasks, err := s.storage.Get(taskId)
	if err != nil {
		// TODO handle error: 404
		WriteJSON(w, http.StatusNotFound, "task not found")
		return
	}

	WriteJSON(w, http.StatusOK, tasks)
}

func (s *Server) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	var payload types.Payload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		WriteJSON(w, http.StatusNotFound, "Error decoding JSON: "+err.Error())
		return
	}

	task := types.NewTask(payload, types.WAITING, time.Now())
	taskId, _ := s.storage.Create(task)
	task.ID = taskId

	go s.pool.AddTask(task)
	WriteJSON(w, http.StatusCreated, task)
}
