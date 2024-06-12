package api

import (
	"encoding/json"
	"net/http"
)

func JsonParse(w http.ResponseWriter, message string) {
	response := map[string]string{"message": message}
	jsonData, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
