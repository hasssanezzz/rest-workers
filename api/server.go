package api

import (
	"github.com/hasssanezzz/rest-workers/storage"
	"github.com/hasssanezzz/rest-workers/types"
	"github.com/hasssanezzz/rest-workers/worker"
)

type Server struct {
	listenAddr string
	storage    *storage.Storage
	pool       *worker.WorkerPool
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
