package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"task-management/models"
)

// In-memory storage for tasks (replace with database in production)
var tasks = make(map[string]models.Task)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	taskList := make([]models.Task, 0, len(tasks))
	for _, task := range tasks {
		taskList = append(taskList, task)
	}
	json.NewEncoder(w).Encode(taskList)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	task, exists := tasks[id]
	if !exists {
		http.Error(w, models.ErrTaskNotFound.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var input models.TaskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := input.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	task := models.Task{
		ID:          id,
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks[id] = task
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	existingTask, exists := tasks[id]
	if !exists {
		http.Error(w, models.ErrTaskNotFound.Error(), http.StatusNotFound)
		return
	}

	var input models.TaskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := input.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	existingTask.Title = input.Title
	existingTask.Description = input.Description
	existingTask.Status = input.Status
	existingTask.UpdatedAt = time.Now()

	tasks[id] = existingTask
	json.NewEncoder(w).Encode(existingTask)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if _, exists := tasks[id]; !exists {
		http.Error(w, models.ErrTaskNotFound.Error(), http.StatusNotFound)
		return
	}

	delete(tasks, id)
	w.WriteHeader(http.StatusNoContent)
} 