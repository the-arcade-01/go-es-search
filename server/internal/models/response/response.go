package models

import (
	"encoding/json"
	"net/http"
)

func ResponseWithJson(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
