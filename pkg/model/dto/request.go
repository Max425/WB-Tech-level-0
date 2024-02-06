package dto

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, response interface{}, statusCode int) {
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
}
