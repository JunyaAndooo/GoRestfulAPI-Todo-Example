package handler

import (
	"net/http"

	"github.com/JunyaAndooo/GoRestfulAPI-Todo-Example/internal/task/entity"
	"github.com/JunyaAndooo/GoRestfulAPI-Todo-Example/internal/task/store"
)

type ListTask struct {
	Store *store.TaskStore
}

type task_output struct {
	Id     entity.TaskId     `json:"id"`
	Title  string            `json:"title"`
	Status entity.TaskStatus `json:"status"`
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks := lt.Store.All()
	rsp := []task_output{}
	for _, t := range tasks {
		rsp = append(rsp, task_output{
			Id:     t.Id,
			Title:  t.Title,
			Status: t.Status,
		})
	}
	RespondJson(ctx, w, rsp, http.StatusOK)
}