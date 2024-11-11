package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"taskmanager/db"
	"taskmanager/models"
	"taskmanager/validation"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	//decode incoming body params to task model
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//validate request
	ValidationErrors := validation.ValidateTask(task)

	if ValidationErrors != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ValidationErrors)
		return
	} else {
		//create task and send created task
		db.DB.Create(&task)
		w.Header().Set("Content-Type", "application/json")
		//encode to json and send task
		json.NewEncoder(w).Encode(task)
	}

}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	// Parse pagination parameters from the query string
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	// Default values for pagination
	page := 1
	limit := 10

	// Convert page and limit to integers
	if pageStr != "" {
		p, err := strconv.Atoi(pageStr)
		if err == nil && p > 0 {
			page = p
		}
	}
	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err == nil && l > 0 {
			limit = l
		}
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Get tasks from the database with pagination
	var tasks []models.Task
	db.DB.Offset(offset).Limit(limit).Find(&tasks)

	// Return tasks as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var task models.Task

	result := db.DB.First(&task, id)
	if result.Error != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var task models.Task
	result := db.DB.First(&task, id)
	if result.Error != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//validate request
	ValidationErrors := validation.ValidateTask(task)

	if ValidationErrors != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ValidationErrors)
		return
	}

	db.DB.Save(&task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var task models.Task
	result := db.DB.First(&task, id)

	if result.Error != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.DB.Delete(&task)
	w.WriteHeader(http.StatusNoContent)
}

func CompleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var task models.Task
	result := db.DB.First(&task, id)
	if result.Error != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//set complete to true of tasks
	task.Completed = true
	db.DB.Save(&task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
