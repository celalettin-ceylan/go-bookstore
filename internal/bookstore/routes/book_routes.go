package routes

import (
	"github.com/celalettin-ceylan/go-bookstore/internal/controllers"
	"github.com/celalettin-ceylan/go-bookstore/internal/repositories"
	"github.com/celalettin-ceylan/go-bookstore/pkg/db"
	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func SetupBookRoutes() *mux.Router {

	bookRepository := repositories.NewBookRepository(db.DB)
	bookController := controllers.NewBookController(bookRepository)

	router.HandleFunc("/books", bookController.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", bookController.GetBookById).Methods("GET")
	router.HandleFunc("/books", bookController.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", bookController.GetAllBooks).Methods("PUT")
	router.HandleFunc("/books/{id}", bookController.DeleteBook).Methods("DELETE")
	return router
}
