package helpers

import (
	"Todo/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func ValidateID(w http.ResponseWriter, idStr string) (int, bool) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return 0, false
	}
	return id, true
}

func ValidateRequest(w http.ResponseWriter, r *http.Request, task *models.Task) bool {
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request to add", http.StatusBadRequest)
		return false
	}
	return true
}
