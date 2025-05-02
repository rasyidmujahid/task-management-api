package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"task-management/handlers"
	"task-management/middleware"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	router := mux.NewRouter()

	// Apply middleware
	router.Use(middleware.CORS)
	router.Use(middleware.ErrorHandler)

	// Health check endpoint
	router.HandleFunc("/health", healthCheck).Methods("GET")

	// Task routes with validation middleware
	taskRouter := router.PathPrefix("/api/tasks").Subrouter()
	taskRouter.Use(middleware.ValidateRequest)

	taskRouter.HandleFunc("", handlers.GetTasks).Methods("GET")
	taskRouter.HandleFunc("", handlers.CreateTask).Methods("POST")
	taskRouter.HandleFunc("/{id}", handlers.GetTask).Methods("GET")
	taskRouter.HandleFunc("/{id}", handlers.UpdateTask).Methods("PUT")
	taskRouter.HandleFunc("/{id}", handlers.DeleteTask).Methods("DELETE")

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
} 