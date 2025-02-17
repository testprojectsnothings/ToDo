package helpers

import (
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
