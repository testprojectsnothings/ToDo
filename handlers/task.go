package handlers

import (
	"Todo/helpers"
	"Todo/models"
	"Todo/storage"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func PostTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if !helpers.ValidateRequest(w, r, &task) {
		return
	}

	task = storage.PostTask(task)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	tasks := storage.GetList()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func PutTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, valid := helpers.ValidateID(w, vars["id"])
	if !valid {
		return
	}

	var updTask models.Task
	if !helpers.ValidateRequest(w, r, &updTask) {
		return
	}

	task, updated := storage.UpdateTasks(id, updTask)
	if !updated {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, valid := helpers.ValidateID(w, vars["id"])
	if !valid {
		return
	}

	if !storage.DeleteTask(id) {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
