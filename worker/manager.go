package worker

import "github.com/hasssanezzz/rest-workers/types"

type AfterFinishFunc func(*types.Task)

type WorkerPool struct {
	Count                int
	Workers              []*Worker
	ReadTaskChan         chan *types.Task
	WriteTaskChan        chan *types.Task
	ProcessFunc          ProcessFunc
	AfterFinishFunc      AfterFinishFunc
	PostStatusUpdateFunc PostStatusUpdateFunc
}

func NewWorkerPool(
	count int,
	payloadChan chan *types.Task,
	resultsChan chan *types.Task,
	processFunc ProcessFunc,
	afterFinishFunc AfterFinishFunc,
	postStatusUpdateFunc PostStatusUpdateFunc,
) *WorkerPool {
	return &WorkerPool{
		Count:                count,
		ReadTaskChan:         payloadChan,
		WriteTaskChan:        resultsChan,
		ProcessFunc:          processFunc,
		AfterFinishFunc:      afterFinishFunc,
		PostStatusUpdateFunc: postStatusUpdateFunc,
	}
}

func (w *WorkerPool) InitiateWorkers() {
	w.Workers = make([]*Worker, w.Count)
	for i := 0; i < w.Count; i++ {
		w.Workers[i] = NewWorker(
			i,
			w.ReadTaskChan,
			w.WriteTaskChan,
			w.PostStatusUpdateFunc,
			w.ProcessFunc,
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
