package main

import (
	"flag"

	"github.com/hasssanezzz/rest-workers/api"
	"github.com/hasssanezzz/rest-workers/storage"
	"github.com/hasssanezzz/rest-workers/types"
	"github.com/hasssanezzz/rest-workers/worker"
)

var payloadChan chan *types.Task
var restulsChan chan *types.Task

func main() {
	listenAddr := flag.String("a", "127.0.0.1:3030", "the listen address in which the server will listen to")
	workerCount := flag.Int("w", 5, "number of workers")
	flag.Parse()

	store := storage.NewStorage()
	payloadChan = make(chan *types.Task, *workerCount)
	restulsChan = make(chan *types.Task, *workerCount)
	pool := worker.NewWorkerPool(
		*workerCount,
		payloadChan,
		restulsChan,
		func(finishedTask *types.Task) {
			// for later use
		},
		func(updatedTask *types.Task) {
			store.UpdateTask(updatedTask)
		},
	)

	server := api.NewServer(*listenAddr, store, pool)
	server.Start()
}
