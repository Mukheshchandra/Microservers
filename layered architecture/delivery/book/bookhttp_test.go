package book

import (
	"GoLang-Interns-2022/entities"
	"GoLang-Interns-2022/service"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetall(t *testing.T) {
	testcases := []struct {
		desc          string
		title         string
		includeAuthor string
		resp          []entities.Book
		err           error
	}{
		{desc: "get all books",
			title:         "",
			includeAuthor: "",
			resp: []entities.Book{{BookID: 1, Title: "2states", Author: entities.Author{AuthorID: 1},
				Publication: "arihant", PublishedDate: "2020/10/23"}, {BookID: 2, Title: "god", Author: entities.Author{AuthorID: 2},
				Publication: "arihant", PublishedDate: "2015/07/03"}},
			err: nil,
		},
		{desc: "data of book with title and including author",
			title:         "god",
			includeAuthor: "true",
			resp: []entities.Book{{BookID: 2, Title: "god", Author: entities.Author{AuthorID: 2, Firstname: "mukhesh", Lastname: "mekala",
				Dob: "2000/10/24", Penname: "zero"}, Publication: "arihant", PublishedDate: "2015/07/03"}},
			err: nil,
		},
	}
	for _, v := range testcases {
		mockCtrl := gomock.NewController(t)

		mockBookService := service.NewMockBookService(mockCtrl)
		mock := New(mockBookService)
		r := httptest.NewRequest("GET", "/books?"+"title="+v.title+"&"+"includeAuthor="+v.includeAuthor, nil)
		w := httptest.NewRecorder()

		mockBookService.EXPECT().GetAllBook(context.TODO(), v.title, v.includeAuthor).Return(v.resp, v.err).AnyTimes()
		mock.GetAllBook(w, r)
		res := w.Result()
		res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Print(err)
		}

		var books []entities.Book

		err = json.Unmarshal(body, &books)
		if err != nil {
			log.Print(err)
		}

		assert.Equal(t, v.resp, books)
	}
}

func TestGetbyID(t *testing.T) {
	testcase := []struct {
		desc          string
		id            int
		resp          entities.Book
		err           error
		expstatuscode int
	}{
		{desc: "valid id",
			id: 1,
			resp: entities.Book{BookID: 1, Title: "2states", Author: entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat",
				Dob: "1980/04/23", Penname: "max"}, Publication: "arihant", PublishedDate: "2020/10/23"},
			err:           nil,
			expstatuscode: http.StatusOK,
		},
		{"invalid id",
			-1,
			entities.Book{},
			errors.New("error"),
			http.StatusBadRequest,
		},
	}
	for _, v := range testcase {
		mockCtrl := gomock.NewController(t)
		mockBookService := service.NewMockBookService(mockCtrl)
		mock := New(mockBookService)

		id := strconv.Itoa(v.id)
		r := httptest.NewRequest("GET", "/books/{id}"+id, nil)
		r = mux.SetURLVars(r, map[string]string{"id": id})
		w := httptest.NewRecorder()

		mockBookService.EXPECT().GetbyID(context.TODO(), v.id).Return(v.resp, v.err).AnyTimes()
		mock.GetbyID(w, r)
		res := w.Result()
		res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Print(err)
		}

		var books entities.Book

		err = json.Unmarshal(body, &books)
		if err != nil {
			log.Print(err)
		}

		assert.Equal(t, v.resp, books)
		assert.Equal(t, v.expstatuscode, res.StatusCode)
	}
}

func TestPostBook(t *testing.T) {
	testcases := []struct {
		desc          string
		req           entities.Book
		resp          entities.Book
		err           error
		expstatuscode int
	}{
		{desc: "details posted for id 1",
			req: entities.Book{BookID: 1, Title: "2states", Author: entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat",
				Dob: "1980/04/23", Penname: "rolex"}, Publication: "arihant", PublishedDate: "2020/10/23"},
			resp: entities.Book{BookID: 1, Title: "2states", Author: entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat",
				Dob: "1980/04/23", Penname: "rolex"}, Publication: "arihant", PublishedDate: "2020/10/23"},
			err:           nil,
			expstatuscode: http.StatusCreated,
		},
		{desc: "details posted for id 2",
			req: entities.Book{BookID: 2, Title: "god", Author: entities.Author{AuthorID: 2, Firstname: "mukhesh", Lastname: "mekala",
				Dob: "2000/10/24", Penname: "zero"}, Publication: "arihant", PublishedDate: "2015/07/03"},
			resp: entities.Book{BookID: 2, Title: "god", Author: entities.Author{AuthorID: 2, Firstname: "mukhesh", Lastname: "mekala",
				Dob: "2000/10/24", Penname: "zero"}, Publication: "arihant", PublishedDate: "2015/07/03"},
			err:           nil,
			expstatuscode: http.StatusCreated,
		},
	}
	for _, v := range testcases {
		mockCtrl := gomock.NewController(t)

		mockBookService := service.NewMockBookService(mockCtrl)
		mock := New(mockBookService)

		data, _ := json.Marshal(v.req)
		r := httptest.NewRequest("POST", "/books", bytes.NewBuffer(data))
		w := httptest.NewRecorder()

		mockBookService.EXPECT().Postbook(context.TODO(), &v.req).Return(v.resp, v.err).AnyTimes()
		mock.Postbook(w, r)
		res := w.Result()
		res.Body.Close()
		assert.Equal(t, v.expstatuscode, res.StatusCode)
	}
}

func TestPutBook(t *testing.T) {
	testcases := []struct {
		desc          string
		reqID         string
		req           entities.Book
		resp          entities.Book
		err           error
		expstatuscode int
	}{
		{desc: "updating details",
			reqID: "1",
			req: entities.Book{BookID: 1, Title: "2states",
				Author:      entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "max"},
				Publication: "arihant", PublishedDate: "2020/10/23"},
			resp: entities.Book{BookID: 1, Title: "2states",
				Author:      entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "max"},
				Publication: "arihant", PublishedDate: "2020/10/23"},
			err:           nil,
			expstatuscode: http.StatusAccepted,
		},
		{desc: "invalid id",
			reqID: "-1",
			req: entities.Book{BookID: -1, Title: "it",
				Author:      entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "max"},
				Publication: "arihant", PublishedDate: "2020/10/23"},
			resp: entities.Book{BookID: -1, Title: "it",
				Author:      entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "max"},
				Publication: "arihant", PublishedDate: "2020/10/23"},
			err:           errors.New("error"),
			expstatuscode: http.StatusBadRequest,
		},
	}
	for _, v := range testcases {
		mockCtrl := gomock.NewController(t)

		mockBookService := service.NewMockBookService(mockCtrl)
		mock := New(mockBookService)

		body, _ := json.Marshal(v.req)
		r := httptest.NewRequest("PUT", "/books/{id}"+v.reqID, bytes.NewBuffer(body))
		r = mux.SetURLVars(r, map[string]string{"id": v.reqID})
		w := httptest.NewRecorder()

		mockBookService.EXPECT().Putbook(context.TODO(), v.req.BookID, &v.req).Return(v.resp, v.err).AnyTimes()
		mock.Putbook(w, r)
		res := w.Result()
		res.Body.Close()
		assert.Equal(t, v.expstatuscode, res.StatusCode)
	}
}

func TestDeleteBook(t *testing.T) {
	testcases := []struct {
		desc               string
		id                 int
		resp               int64
		err                error
		expectedstatuscode int
	}{
		{"success:valid id", 1, 1, nil, http.StatusNoContent},
		{"failure:invalid id", -9, 0, errors.New("error"), http.StatusBadRequest},
	}
	for _, v := range testcases {
		mockCtrl := gomock.NewController(t)

		mockBookService := service.NewMockBookService(mockCtrl)

		mock := New(mockBookService)

		id := strconv.Itoa(v.id)
		r := httptest.NewRequest("DELETE", "/books/{id}"+id, nil)
		r = mux.SetURLVars(r, map[string]string{"id": id})
		w := httptest.NewRecorder()

		mockBookService.EXPECT().Deletebook(context.TODO(), v.id).Return(v.resp, v.err).AnyTimes()
		mock.Deletebook(w, r)
		res := w.Result()
		res.Body.Close()
		assert.Equal(t, v.expectedstatuscode, res.StatusCode)
	}
}
