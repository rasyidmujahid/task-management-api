package main

import (
	"log"
	"net/http"
	"os"

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

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
} 