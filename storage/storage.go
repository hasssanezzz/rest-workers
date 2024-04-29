package storage

import (
	"fmt"

	"github.com/hasssanezzz/rest-workers/types"
)

var lastId = 0

type Storage struct {
	Tasks map[int]*types.Task
}

func NewStorage() *Storage {
	return &Storage{
		Tasks: make(map[int]*types.Task),
	}
}

func (s *Storage) Get(id int) (*types.Task, error) {
	if task, ok := s.Tasks[id]; ok {
		return task, nil
	} else {
		return nil, fmt.Errorf("task with id (%d) not found", id)
	}
}

func (s *Storage) List() []*types.Task {
	tasks := make([]*types.Task, 0, len(s.Tasks))
	for _, value := range s.Tasks {
		tasks = append(tasks, value)
	}
	return tasks
}

func (s *Storage) Create(task *types.Task) (int, error) {
	task.ID = lastId
	s.Tasks[lastId] = task
	lastId++
	return task.ID, nil
}

func (s *Storage) UpdateTask(updatedTask *types.Task) error {
	if _, ok := s.Tasks[updatedTask.ID]; !ok {
		return fmt.Errorf("task with id (%d) not found", updatedTask.ID)
	}

	delete(s.Tasks, updatedTask.ID)
	s.Tasks[updatedTask.ID] = updatedTask

	return nil
}

func (s *Storage) Delete(id int) {
	delete(s.Tasks, id)
}
