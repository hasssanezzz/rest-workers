package api

import (
	"encoding/json"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/hasssanezzz/rest-workers/storage"
	"github.com/hasssanezzz/rest-workers/types"
	"github.com/hasssanezzz/rest-workers/worker"
)

var payloadChan chan *types.Task = make(chan *types.Task)
var restulsChan chan *types.Task = make(chan *types.Task)

type Server struct {
	listenAddr string
	storage    *storage.Storage
	pool       *worker.WorkerPool
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func NewServer(listenAddr string, workerCount int) *Server {
	localStorage := storage.NewStorage()

	pool := worker.NewWorkerPool(
		5,
		payloadChan,
		restulsChan,
		func(finishedTask *types.Task) {

		},
		func(updatedTask *types.Task) {
			localStorage.UpdateTask(updatedTask)
		},
	)

	go pool.InitiateWorkers()

	return &Server{
		listenAddr: listenAddr,
		storage:    localStorage,
		pool:       pool,
	}
}

func (s *Server) Start() error {
	router := mux.NewRouter()

	router.HandleFunc("/api/v0/task", s.handleListTasks).Methods("GET")
	router.HandleFunc("/api/v0/task/{id}", s.handleGetTask).Methods("GET")
	router.HandleFunc("/api/v0/task", s.handleCreateTask).Methods("POST")

	return http.ListenAndServe(s.listenAddr, router)
}

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
	type RequestBody struct {
		Value string `json:"value"`
	}

	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil || requestBody.Value == "" {
		WriteJSON(w, http.StatusBadRequest, "bad request body")
		return
	}

	bigInt, ok := big.NewInt(0).SetString(requestBody.Value, 10)
	if !ok {
		WriteJSON(w, http.StatusBadRequest, "can not parse the provided number")
		return
	}

	payload := types.Payload{Number: bigInt}
	task := types.NewTask(payload, types.WAITING, time.Now())
	taskId, _ := s.storage.Create(task)
	task.ID = taskId

	go s.pool.AddTask(task)
	WriteJSON(w, http.StatusCreated, task)
}
