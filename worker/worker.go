package worker

import (
	"log"
	"time"

	"github.com/hasssanezzz/rest-workers/types"
)

type Status string

type PostStatusUpdateFunc func(*types.Task)
type ProcessFunc func(*types.Payload) *types.Result

const (
	WAITING Status = "Waiting"
	WORKING Status = "Working"
)

type Worker struct {
	Index                int
	Status               Status
	PostStatusUpdateFunc PostStatusUpdateFunc
	ProcessFunc          ProcessFunc

	readch  <-chan *types.Task
	writech chan<- *types.Task
}

func NewWorker(
	index int,
	readch <-chan *types.Task,
	writech chan<- *types.Task,
	postStatusUpdateFunc PostStatusUpdateFunc,
	processFunc ProcessFunc) *Worker {

	return &Worker{
		Index:                index,
		Status:               WAITING,
		PostStatusUpdateFunc: postStatusUpdateFunc,
		ProcessFunc:          processFunc,
		readch:               readch,
		writech:              writech,
	}
}

func (w *Worker) RunAndListen() {
	for task := range w.readch {

		// update worker status to "working"
		w.Status = WORKING
		// update the start date
		task.StartedAt = time.Now()
		// update task status to "working"
		task.Status = types.WORKING
		// apply the update to the storage
		w.PostStatusUpdateFunc(task)

		log.Printf("Worker:%d :: STARTED\t %d\n", w.Index, task.ID)

		// result := w.compute(&task.Payload)
		result := w.ProcessFunc(&task.Payload)

		// after the worker finished working, it is now free
		// and waiting for new taks so we upate
		// the status to "waiting"
		w.Status = WAITING
		// update the task results
		task.Result = *result
		// update the finish data
		task.FinishedAt = time.Now()
		// update task status to "finished"
		task.Status = types.FINISHED
		// apply the update to the storage
		w.PostStatusUpdateFunc(task)

		log.Printf("Worker:%d :: FINISHED\t %d\n", w.Index, task.ID)

		// write the result
		w.writech <- task
	}
}
