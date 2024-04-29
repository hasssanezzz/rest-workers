package worker

import "github.com/hasssanezzz/rest-workers/types"

type AfterFinishFunc func(*types.Task)

type WorkerPool struct {
	Count                int
	Workers              []*PrimeAnalyzerWorker
	ReadTaskChan         chan *types.Task
	WriteTaskChan        chan *types.Task
	AfterFinishFunc      AfterFinishFunc
	PostStatusUpdateFunc PostStatusUpdateFunc
}

func NewWorkerPool(
	count int,
	payloadChan chan *types.Task,
	resultsChan chan *types.Task,
	afterFinishFunc AfterFinishFunc,
	postStatusUpdateFunc PostStatusUpdateFunc,
) *WorkerPool {

	return &WorkerPool{
		Count:                count,
		ReadTaskChan:         payloadChan,
		WriteTaskChan:        resultsChan,
		AfterFinishFunc:      afterFinishFunc,
		PostStatusUpdateFunc: postStatusUpdateFunc,
	}
}

func (w *WorkerPool) InitiateWorkers() {
	w.Workers = make([]*PrimeAnalyzerWorker, w.Count)
	for i := 0; i < w.Count; i++ {
		w.Workers[i] = NewPrimeAnalyzerWorker(
			w.ReadTaskChan,
			w.WriteTaskChan,
			w.PostStatusUpdateFunc,
			i,
		)

		go w.Workers[i].RunAndListen()
		go w.ResultChanListen()
	}
}

func (w *WorkerPool) AddTask(task *types.Task) {
	w.ReadTaskChan <- task
}

func (w *WorkerPool) ResultChanListen() {
	for finishedTask := range w.WriteTaskChan {
		w.AfterFinishFunc(finishedTask)
	}
}
