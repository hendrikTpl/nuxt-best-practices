package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	if id == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
		return
	}

	// getById
	for _, book := range books {
		if book.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

// Add
func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	if book.Author == "" || book.Title == "" {
		http.Error(w, "Please provide a title and an author", http.StatusBadRequest)
		return
	}

	book.ID = len(books) + 1
	books = append(books, book)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func main() {
	books = append(books, Book{ID: 1, Title: "The Go Programming Language", Author: "Alan A. A. Donovan"})
	books = append(books, Book{ID: 2, Title: "Learning Go", Author: "Jon Bodner"})
	r := mux.NewRouter()
	// Endpoint
	r.HandleFunc("/books/", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")

	fmt.Println("Server is running on port 5000...")
	log.Fatal(http.ListenAndServe(":5000", r))
}
