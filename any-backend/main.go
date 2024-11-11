package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"taskmanager/db"
	"taskmanager/handlers"
)

func main() {
	//load environment variables from .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
		return
	}

	//connect to database
	db.Connect()
	db.Migrate()

	//setup router
	router := mux.NewRouter()

	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")
	router.HandleFunc("/complete/{id}", handlers.CompleteTask).Methods("GET")

	//start the server
	log.Println("Server is running on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
