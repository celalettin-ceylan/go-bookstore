package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/celalettin-ceylan/go-bookstore/internal/models"
	"github.com/celalettin-ceylan/go-bookstore/internal/repositories"
	"github.com/gorilla/mux"
)

type BookController struct {
	Repo *repositories.BookRepository
}

func NewBookController(bookrepository *repositories.BookRepository) *BookController {
	return &BookController{Repo: bookrepository}
}

func (bc *BookController) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := bc.Repo.GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}

func (bc *BookController) GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err := bc.Repo.GetBookById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func (bc *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdBook, err := bc.Repo.CreateBook(&book)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdBook)
}

func (bc *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book.ID = uint(id)
	updatedBook, err := bc.Repo.UpdateBook(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedBook)
}

func (bc *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err2 := bc.Repo.DeleteBook(uint(id))
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Deleted Book")
}
