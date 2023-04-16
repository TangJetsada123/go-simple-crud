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

type Books struct {
	ID      string   `json:"id"`
	Isbn    string   `json:"isbn"`
	Title   string   `json:"title"`
	Article *Article `json:"article"`
}

type Article struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Books

func helloFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello")
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Books
	_ = json.NewDecoder(r.Body).Decode(&books)
	book.ID = strconv.Itoa(rand.Intn(100000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

func updateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Books
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
		}
	}
}

func deleteBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	r := mux.NewRouter()

	books = append(books, Books{ID: "1", Isbn: "4620", Title: "The lord of the ring", Article: &Article{Firstname: "Peter", Lastname: "Parker"}})
	books = append(books, Books{ID: "2", Isbn: "4620", Title: "Harry Potter", Article: &Article{Firstname: "JK", Lastname: "Rollings"}})
	r.HandleFunc("/", helloFunc).Methods("GET")
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBooks).Methods("UPDATE")
	r.HandleFunc("/books/{id}", deleteBooks).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
