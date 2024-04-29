package worker

import (
	"log"
	"math/big"
	"time"

	"github.com/hasssanezzz/rest-workers/types"
)

type Status string

type PostStatusUpdateFunc func(*types.Task)

const (
	WAITING Status = "Waiting"
	WORKING Status = "Working"
)

type PrimeAnalyzerWorker struct {
	Index                int
	Status               Status
	PostStatusUpdateFunc PostStatusUpdateFunc

	readch  <-chan *types.Task
	writech chan<- *types.Task
}

func NewPrimeAnalyzerWorker(
	readch <-chan *types.Task,
	writech chan<- *types.Task,
	postStatusUpdateFunc PostStatusUpdateFunc,
	index int) *PrimeAnalyzerWorker {

	return &PrimeAnalyzerWorker{
		Index:                index,
		Status:               WAITING,
		PostStatusUpdateFunc: postStatusUpdateFunc,
		readch:               readch,
		writech:              writech,
	}
}

func (w *PrimeAnalyzerWorker) RunAndListen() {
	for task := range w.readch {

		w.Status = WORKING
		task.StartedAt = time.Now() // here we have reflect the status immediately
		task.Status = types.WORKING
		w.PostStatusUpdateFunc(task)
		log.Printf("Worker:%d :: STARTED\t %d\n", w.Index, task.ID)

		result := w.compute(*task.Payload.Number)

		w.Status = WAITING
		task.Result.Result = result
		task.FinishedAt = time.Now()
		task.Status = types.FINISHED
		w.PostStatusUpdateFunc(task)
		log.Printf("Worker:%d :: FINISHED\t %d\n", w.Index, task.ID)

		// write the result
		w.writech <- task
	}
}

func (w *PrimeAnalyzerWorker) compute(x big.Int) bool {
	if x.Cmp(big.NewInt(1)) <= 0 || x.Cmp(big.NewInt(2)) > 0 && x.Bit(0) == 0 {
		return false
	}

	itr := new(big.Int).SetInt64(2)
	sqrtX := new(big.Int).Sqrt(&x)
	for itr.Cmp(sqrtX) <= 0 {
		if new(big.Int).Mod(&x, itr).Cmp(big.NewInt(0)) == 0 {
			return false
		}
		itr.Add(itr, big.NewInt(1))
	}

	return true
}
