package store

import (
	"errors"

	"github.com/JunyaAndooo/GoRestfulAPI-Todo-Example/internal/task/entity"
)

var (
	Tasks       = &TaskStore{Tasks: map[entity.TaskId]*entity.TaskEntity{}}
	ErrNotFound = errors.New("not found")
)

// 動作確認用
type TaskStore struct {
	LastId entity.TaskId
	Tasks  map[entity.TaskId]*entity.TaskEntity
}

func (ts *TaskStore) Add(t *entity.TaskEntity) (int, error) {
	ts.LastId++
	t.Id = ts.LastId
	ts.Tasks[t.Id] = t
	return int(t.Id), nil
}

func (ts *TaskStore) All() entity.Tasks {
	tasks := make([]*entity.TaskEntity, len(ts.Tasks))
	for i, t := range ts.Tasks {
		tasks[i-1] = t
	}

	return tasks
}
