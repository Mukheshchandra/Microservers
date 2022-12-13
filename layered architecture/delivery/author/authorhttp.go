package author

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
	service service.AuthorService
}

func New(s service.AuthorService) Handler {
	return Handler{service: s}
}

func (a Handler) Postauthor(w http.ResponseWriter, r *http.Request) {
	var author entities.Author

	body, _ := io.ReadAll(r.Body)

	err := json.Unmarshal(body, &author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	ctx := r.Context()

	res, err := a.service.Postauthor(ctx, author)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	auth, err := json.Marshal(res)
	if err != nil {
		log.Print(err)
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(auth)
}

func (a Handler) Putauthor(w http.ResponseWriter, r *http.Request) {
	var author entities.Author

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	body, _ := io.ReadAll(r.Body)

	err = json.Unmarshal(body, &author)
	if err != nil {
		log.Print(err)
	}

	ctx := context.TODO()

	res, err := a.service.Putauthor(ctx, id, author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	auth, err := json.Marshal(res)
	if err != nil {
		log.Print(err)
	}

	w.WriteHeader(http.StatusAccepted)
	_, _ = w.Write(auth)
}

func (a Handler) Deleteauthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	ctx := context.TODO()

	_, err = a.service.Deleteauthor(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
