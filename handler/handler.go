package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrResponse struct {
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}

func RespondJson(ctx context.Context, w http.ResponseWriter, body any, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		errStatus := http.StatusInternalServerError
		w.WriteHeader(errStatus)
		rsp := ErrResponse{
			Message: http.StatusText(errStatus),
		}
		err = json.NewEncoder(w).Encode(rsp)
		if err != nil {
			fmt.Printf("write error response error: %v", err)
		}

		return
	}

	w.WriteHeader(status)
	_, err = fmt.Fprintf(w, "%s", bodyBytes)
	if err != nil {
		fmt.Printf("write response error: %v", err)
	}
}
