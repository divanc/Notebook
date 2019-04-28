package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// books var
var books []Book

// Get all books
func getBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(books)
}

// Get THE book
func getBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request) // Get params

	// Find book with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	json.NewEncoder(writer).Encode(&Book{})
}

// Create a new book
func createBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var book Book

	decoder := json.NewDecoder(request.Body)
	error := decoder.Decode(&book)

	if error != nil {
		panic(error)
	}

	fmt.Println(book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID — Not safe
	books = append(books, book)

	json.NewEncoder(writer).Encode(book)
}

// Update the book
func updateBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)

			var book Book
			_ = json.NewDecoder(request.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			fmt.Println("Book updated: ", book)
			return
		}
	}
	fmt.Println("Book updated")
	json.NewEncoder(writer).Encode(books)
}

// Delete book
func deleteBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	fmt.Println("Book deleted")
	json.NewEncoder(writer).Encode(books)
}

func main() {
	// Init Router
	router := mux.NewRouter()

	// Mock data @todo — implement DB
	books = append(books, Book{ID: "1", Isbn: "44534243", Title: "Happy Day",
		Author: &Author{Firstname: "Louis", Lastname: "Down"}})
	books = append(books, Book{ID: "2", Isbn: "44234243", Title: "Long Life",
		Author: &Author{Firstname: "Verno", Lastname: "Banani"}})

	// Route handlers / Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
