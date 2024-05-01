package api

import (
	"github.com/hasssanezzz/rest-workers/storage"
	"github.com/hasssanezzz/rest-workers/worker"
)

type Server struct {
	listenAddr string
	storage    *storage.Storage
	pool       *worker.WorkerPool
}

func NewServer(listenAddr string, store *storage.Storage, pool *worker.WorkerPool) *Server {
	go pool.InitiateWorkers()

	return &Server{
		listenAddr: listenAddr,
		storage:    store,
		pool:       pool,
	}
}
