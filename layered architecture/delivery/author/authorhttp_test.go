package author

import (
	"GoLang-Interns-2022/entities"
	"GoLang-Interns-2022/service"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestPostauthor(t *testing.T) {
	testcases := []struct {
		desc          string
		req           entities.Author
		resp          entities.Author
		err           error
		expstatuscode int
	}{
		{"valid author",
			entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "max"},
			entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "max"},
			nil,
			http.StatusCreated},
		{desc: "invalid id",
			req:           entities.Author{AuthorID: -1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "rolex"},
			resp:          entities.Author{AuthorID: -1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "rolex"},
			err:           errors.New("error"),
			expstatuscode: http.StatusBadRequest},
	}
	for _, v := range testcases {
		mockCtrl := gomock.NewController(t)

		mockAuthorService := service.NewMockAuthorService(mockCtrl)
		mock := New(mockAuthorService)

		data, err := json.Marshal(v.req)
		if err != nil {
			log.Print(err)
		}

		r := httptest.NewRequest("POST", "/author", bytes.NewBuffer(data))
		w := httptest.NewRecorder()

		mockAuthorService.EXPECT().Postauthor(context.TODO(), v.req).Return(v.resp, v.err).AnyTimes()
		mock.Postauthor(w, r)
		res := w.Result()
		res.Body.Close()
		assert.Equal(t, v.expstatuscode, res.StatusCode)
	}
}

func TestPutauthor(t *testing.T) {
	testcases := []struct {
		desc          string
		reqID         string
		req           entities.Author
		resp          entities.Author
		err           error
		expstatuscode int
	}{
		{desc: "valid author",
			reqID:         "1",
			req:           entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "rolex"},
			resp:          entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "rolex"},
			err:           nil,
			expstatuscode: http.StatusAccepted,
		},
		{desc: "invalid id",
			reqID:         "-1",
			req:           entities.Author{AuthorID: -1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "rolex"},
			resp:          entities.Author{AuthorID: -1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "rolex"},
			err:           errors.New("error"),
			expstatuscode: http.StatusBadRequest,
		},
	}
	for _, v := range testcases {
		mockCtrl := gomock.NewController(t)

		mockAuthorService := service.NewMockAuthorService(mockCtrl)
		mock := New(mockAuthorService)

		data, err := json.Marshal(v.req)
		if err != nil {
			log.Print(err)
		}

		r := httptest.NewRequest("POST", "/author/{id}"+v.reqID, bytes.NewBuffer(data))
		r = mux.SetURLVars(r, map[string]string{"id": v.reqID})
		w := httptest.NewRecorder()

		mockAuthorService.EXPECT().Putauthor(context.TODO(), v.req.AuthorID, v.req).Return(v.resp, v.err).AnyTimes()
		mock.Putauthor(w, r)
		res := w.Result()
		res.Body.Close()
		assert.Equal(t, v.expstatuscode, res.StatusCode)
	}
}

func TestDeleteauthor(t *testing.T) {
	testcases := []struct {
		desc               string
		id                 int
		resp               int64
		err                error
		expectedstatuscode int
	}{
		{"success:valid id", 1, 1, nil, http.StatusNoContent},
		{"failure:invalid id", -4, 0, errors.New("error"), http.StatusBadRequest},
	}
	for _, v := range testcases {
		mockCtrl := gomock.NewController(t)

		mockAuthorService := service.NewMockAuthorService(mockCtrl)
		mock := New(mockAuthorService)

		id := strconv.Itoa(v.id)
		r := httptest.NewRequest("DELETE", "/author/{id}"+id, nil)
		r = mux.SetURLVars(r, map[string]string{"id": id})
		w := httptest.NewRecorder()

		mockAuthorService.EXPECT().Deleteauthor(context.TODO(), v.id).Return(v.resp, v.err).AnyTimes()
		mock.Deleteauthor(w, r)
		res := w.Result()
		res.Body.Close()
		assert.Equal(t, v.expectedstatuscode, res.StatusCode)
	}
}
