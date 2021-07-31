package helper

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type PaginationResponse struct {
	Page int         `json:"page,omitempty"`
	Size int         `json:"size,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func (*Response) Send(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
