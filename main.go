package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"task-management/handlers"
)

func main() {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/api/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/api/tasks", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", handlers.GetTask).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	router.HandleFunc("/api/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
} 