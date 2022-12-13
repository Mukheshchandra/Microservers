package book

import (
	"GoLang-Interns-2022/entities"
	"GoLang-Interns-2022/service"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	service service.BookService
}

func New(s service.BookService) Handler {
	return Handler{service: s}
}

func (b Handler) GetAllBook(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	includeAuthor := r.URL.Query().Get("includeAuthor")

	ctx := r.Context()

	books, _ := b.service.GetAllBook(ctx, title, includeAuthor)

	data, err := json.Marshal(books)
	if err != nil {
		log.Print(err)
	}

	_, _ = w.Write(data)
}
func (b Handler) GetbyID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil || id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := context.TODO()

	book, err := b.service.GetbyID(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	data, err := json.Marshal(book)
	if err != nil {
		_, _ = w.Write([]byte("book cannot be encoded"))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func (b Handler) Postbook(w http.ResponseWriter, r *http.Request) {
	var book entities.Book

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	books, err := b.service.Postbook(ctx, &book)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	data, err := json.Marshal(books)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(data)
}

func (b Handler) Putbook(w http.ResponseWriter, r *http.Request) {
	var book entities.Book

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	body, _ := io.ReadAll(r.Body)

	err = json.Unmarshal(body, &book)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	ctx := context.TODO()

	_, err = b.service.Putbook(ctx, id, &book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	_, _ = w.Write(body)
}

func (b Handler) Deletebook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil || id <= 0 {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	ctx := context.TODO()

	_, err = b.service.Deletebook(ctx, id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
