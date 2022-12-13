package main

import (
	"log"
	"net/http"

	"GoLang-Interns-2022/datastore/author"
	"GoLang-Interns-2022/datastore/book"
	authorHdlr "GoLang-Interns-2022/delivery/author"
	bookHdlr "GoLang-Interns-2022/delivery/book"
	"GoLang-Interns-2022/driver"
	authorSvc "GoLang-Interns-2022/service/author"
	bookSvc "GoLang-Interns-2022/service/book"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	db := driver.Connection()
	authorStore := author.New(db)
	authorService := authorSvc.New(authorStore)
	authorHandler := authorHdlr.New(authorService)

	bookStore := book.New(db)
	bookService := bookSvc.New(bookStore, authorStore)
	bookHandler := bookHdlr.New(bookService)

	r.HandleFunc("/author", authorHandler.Postauthor).Methods("POST")
	r.HandleFunc("/author/{id}", authorHandler.Putauthor).Methods("PUT")
	r.HandleFunc("/author/{id}", authorHandler.Deleteauthor).Methods("DELETE")

	r.HandleFunc("/books", bookHandler.GetAllBook).Methods("GET")
	r.HandleFunc("/books/{id}", bookHandler.GetbyID).Methods("GET")
	r.HandleFunc("/books", bookHandler.Postbook).Methods("POST")
	r.HandleFunc("/books/{id}", bookHandler.Putbook).Methods("PUT")
	r.HandleFunc("/books/{id}", bookHandler.Deletebook).Methods("DELETE")

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Print(err)
	}
}
