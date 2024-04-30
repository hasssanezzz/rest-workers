package api

import (
	"encoding/json"
	"net/http"

	"github.com/hasssanezzz/rest-workers/storage"
	"github.com/hasssanezzz/rest-workers/types"
	"github.com/hasssanezzz/rest-workers/worker"
)

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
	payloadChan := make(chan *types.Task, workerCount)
	restulsChan := make(chan *types.Task, workerCount)

	localStorage := storage.NewStorage()

	pool := worker.NewWorkerPool(
		workerCount,
		payloadChan,
		restulsChan,
		func(finishedTask *types.Task) {
			// for later use
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
