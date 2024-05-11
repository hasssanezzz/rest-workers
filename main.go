package main

import (
	"flag"
	"math/big"

	"github.com/hasssanezzz/rest-workers/api"
	"github.com/hasssanezzz/rest-workers/storage"
	"github.com/hasssanezzz/rest-workers/types"
	"github.com/hasssanezzz/rest-workers/worker"
)

var payloadChan chan *types.Task
var restulsChan chan *types.Task

func ProcessFunc(payload *types.Payload) *types.Result {
	x := *payload.Number

	if x.Cmp(big.NewInt(1)) <= 0 || x.Cmp(big.NewInt(2)) > 0 && x.Bit(0) == 0 {
		return &types.Result{
			Result: false,
		}
	}

	itr := new(big.Int).SetInt64(2)
	sqrtX := new(big.Int).Sqrt(&x)
	for itr.Cmp(sqrtX) <= 0 {
		if new(big.Int).Mod(&x, itr).Cmp(big.NewInt(0)) == 0 {
			return &types.Result{
				Result: false,
			}
		}
		itr.Add(itr, big.NewInt(1))
	}

	return &types.Result{
		Result: true,
	}
}

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
		ProcessFunc,
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
