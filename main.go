package main

import (
	"Todo/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/task/list", handlers.GetList).Methods("GET")
	router.HandleFunc("/task/update/{id}", handlers.PutTasks).Methods("PUT")
	router.HandleFunc("/task/create", handlers.PostTask).Methods("POST")
	router.HandleFunc("/task/delete/{id}", handlers.DeleteTask).Methods("DELETE")

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", router)
}
