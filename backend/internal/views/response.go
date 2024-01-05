package views

import (
	"encoding/json"
	"net/http"
)

func SendJSONResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
}

func SendResponse(w http.ResponseWriter, data any) {
	SendJSONResponse(w, http.StatusOK, data)
}

func SendErrorMsg(w http.ResponseWriter, data any) {
	SendJSONResponse(w, http.StatusBadRequest, data)
}
