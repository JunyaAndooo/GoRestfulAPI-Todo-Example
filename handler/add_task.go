package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JunyaAndooo/GoRestfulAPI-Todo-Example/internal/task/entity"
	"github.com/JunyaAndooo/GoRestfulAPI-Todo-Example/internal/task/store"
	"github.com/go-playground/validator/v10"
)

type AddTask struct {
	Store     *store.TaskStore
	Validator *validator.Validate
}

func (at *AddTask) ServerHttp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var b struct {
		Title string `json:"title" validate:"required"`
	}

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		RespondJson(
			ctx, 
			w, 
			&ErrResponse{
				Message: err.Error(),
			},
			http.StatusInternalServerError,
		)

		return
	}

	err = at.Validator.Struct(b)
	if err != nil {
		RespondJson(
			ctx, 
			w, 
			&ErrResponse{
				Message: err.Error(),
			},
			http.StatusBadRequest,
		)

		return
	}

	t := &entity.TaskEntity{
		Title: b.Title,
		Status: entity.TaskStatusTodo,
		CreatedAt: time.Now(),
	}

	id, err := store.Tasks.Add(t)
	if err != nil {
		RespondJson(
			ctx, 
			w, 
			&ErrResponse{
				Message: err.Error(),
			},
			http.StatusInternalServerError,
		)

		return
	}

	rsp := struct {
		Id int `json:"id"`
	}{Id:id}

	RespondJson(ctx, w, rsp, http.StatusOK)
}