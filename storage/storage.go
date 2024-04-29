package storage

import (
	"fmt"
	"time"

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

func (s *Storage) UpdateTaskStatus(id int, status types.Status) error {
	if _, ok := s.Tasks[id]; ok {
		s.Tasks[id].Status = status
	}

	return nil
}

func (s *Storage) PlaceTask(id int, placedAt time.Time) error {
	if _, ok := s.Tasks[id]; ok {
		s.Tasks[id].PlacedAt = placedAt
	}

	return nil
}

func (s *Storage) StartTask(id int, startedaAt time.Time) error {
	if _, ok := s.Tasks[id]; ok {
		s.Tasks[id].StartedAt = startedaAt
	}

	return nil
}

func (s *Storage) FinishTask(id int, result *types.Result, finishedAt time.Time) error {
	if _, ok := s.Tasks[id]; ok {
		s.Tasks[id].Result = *result // TODO does this have to be a pointer?
		s.Tasks[id].FinishedAt = finishedAt
	}

	return nil
}

func (s *Storage) Delete(id int) {
	delete(s.Tasks, id)
}
