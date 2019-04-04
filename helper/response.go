package helper

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	HttpCode int         `json:"httpCode"`
	Datas     interface{} `json:"data"`
}

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
