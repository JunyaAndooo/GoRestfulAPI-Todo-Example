package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JunyaAndooo/GoRestfulAPI-Todo-Example/internal/task/entity"
	"github.com/JunyaAndooo/GoRestfulAPI-Todo-Example/internal/task/store"
	"github.com/JunyaAndooo/GoRestfulAPI-Todo-Example/testutil"
	"github.com/go-playground/validator/v10"
)

func TestAddTask(t *testing.T) {
	type want struct {
		status  int
		rspFile string
	}

	tests := map[string]struct {
		reqFile string
		want    want
	}{
		"ok": {
			reqFile: "testdata/add_task/ok_req.json.golden",
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/add_task/ok_rsp.json.golden",
			},
		},
		"badRequest": {
			reqFile: "testdata/add_task/bad_req.json.golden",
			want: want{
				status:  http.StatusBadRequest,
				rspFile: "testdata/add_task/bad_rsp.json.golden",
			},
		},
	}

	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPost,
				"/tasks",
				bytes.NewReader(testutil.LoadFile(t, tt.reqFile)),
			)

			sut := AddTask{Store: &store.TaskStore{
				Tasks: map[entity.TaskId]*entity.TaskEntity{},
			}, Validator: validator.New()}
			sut.ServerHttp(w, r)

			resp := w.Result()
			testutil.AssertResponse(t,
				resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile),
			)
		})
	}
}