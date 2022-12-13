package book

import (
	"context"
	"errors"
	"testing"

	"GoLang-Interns-2022/datastore"
	"GoLang-Interns-2022/entities"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetallbook(t *testing.T) {
	author := entities.Author{AuthorID: 2, Firstname: "mukhesh", Lastname: "mekala", Dob: "2000/10/24", Penname: "zero"}

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
				Publication: "arihant", PublishedDate: "2020/10/23"},
				{BookID: 2, Title: "god", Author: entities.Author{AuthorID: 2}, Publication: "arihant", PublishedDate: "2015/07/03"}},
			err: nil,
		},
		{desc: "data of book with title and including author",
			title:         "god",
			includeAuthor: "true",
			resp: []entities.Book{{BookID: 2, Title: "god", Author: entities.Author{AuthorID: 2, Firstname: "mukhesh", Lastname: "mekala",
				Dob: "2000/10/24", Penname: "zero"}, Publication: "arihant", PublishedDate: "2015/07/03"}},
			err: nil},
	}
	for _, v := range testcases {
		mockCtrl := gomock.NewController(t)

		mockBookStore := datastore.NewMockBook(mockCtrl)
		mockAuthorStore := datastore.NewMockAuthor(mockCtrl)
		mock := New(mockBookStore, mockAuthorStore)

		mockAuthorStore.EXPECT().IncludeAuthor(context.TODO(), author.AuthorID).Return(author, v.err).AnyTimes()
		mockBookStore.EXPECT().GetAllBook(context.TODO()).Return(v.resp, v.err).AnyTimes()
		mockBookStore.EXPECT().GetBookbyTitle(context.TODO(), v.title).Return(v.resp, v.err).AnyTimes()

		res, err := mock.GetAllBook(context.TODO(), v.title, v.includeAuthor)

		assert.Equal(t, v.resp, res)
		assert.Equal(t, v.err, err)
	}
}

func TestGetbyID(t *testing.T) {
	testcases := []struct {
		desc string
		req  int
		resp entities.Book
		err  error
	}{
		{desc: "fetching details of book by id",
			req: 1,
			resp: entities.Book{BookID: 1, Title: "2states", Author: entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat",
				Dob: "1980/04/23", Penname: "max"}, Publication: "arihant", PublishedDate: "2020/10/23"},
			err: nil,
		},
	}
	for _, v := range testcases {
		mockCtrl := gomock.NewController(t)

		mockBookStore := datastore.NewMockBook(mockCtrl)
		mockAuthorStore := datastore.NewMockAuthor(mockCtrl)
		mock := New(mockBookStore, mockAuthorStore)

		mockBookStore.EXPECT().GetbyID(context.TODO(), v.req).Return(v.resp, v.err).AnyTimes()
		res, err := mock.GetbyID(context.TODO(), v.req)

		assert.Equal(t, v.err, err)
		assert.Equal(t, v.resp, res)
	}
}

func TestPostbook(t *testing.T) {
	testcases := []struct {
		desc     string
		req      entities.Book
		response entities.Book
		err      error
		err1     error
	}{
		{desc: "valid book",
			req: entities.Book{BookID: 1, Title: "2states", Author: entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat",
				Dob: "1980/04/23", Penname: "rolex"}, Publication: "arihant", PublishedDate: "2020/10/23"},
			response: entities.Book{BookID: 1, Title: "2states", Author: entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat",
				Dob: "1980/04/23", Penname: "rolex"}, Publication: "arihant", PublishedDate: "2020/10/23"},
			err:  nil,
			err1: nil,
		},
		{
			desc: "valid book",
			req: entities.Book{BookID: 2, Title: "god", Author: entities.Author{AuthorID: 2, Firstname: "mukhesh", Lastname: "mekala",
				Dob: "2000/10/24", Penname: "zero"}, Publication: "arihant", PublishedDate: "2015/07/03"},
			response: entities.Book{BookID: 2, Title: "god", Author: entities.Author{AuthorID: 2, Firstname: "mukhesh", Lastname: "mekala",
				Dob: "2000/10/24", Penname: "zero"}, Publication: "arihant", PublishedDate: "2015/07/03"},
			err:  nil,
			err1: nil,
		},
		{
			desc: "invalid book id",
			req: entities.Book{BookID: -2, Title: "god", Author: entities.Author{AuthorID: 2, Firstname: "mukhesh", Lastname: "mekala",
				Dob: "2000/10/24", Penname: "zero"}, Publication: "arihant", PublishedDate: "2015/07/03"},
			response: entities.Book{},
			err:      errors.New("invalid credentials"),
			err1:     nil,
		},
	}
	for _, v := range testcases {
		mockCtrl := gomock.NewController(t)

		mockAuthorStore := datastore.NewMockAuthor(mockCtrl)
		mockBookStore := datastore.NewMockBook(mockCtrl)
		mock := New(mockBookStore, mockAuthorStore)

		mockBookStore.EXPECT().Postbook(context.TODO(), &v.req).Return(v.response, v.err).AnyTimes()
		mockAuthorStore.EXPECT().IncludeAuthor(context.TODO(), v.req.Author.AuthorID).Return(v.req.Author, v.err1).AnyTimes()

		res, err := mock.Postbook(context.TODO(), &v.req)

		assert.Equal(t, v.err, err)
		assert.Equal(t, v.response, res)
	}
}

func TestPutbook(t *testing.T) {
	testcases := []struct {
		desc  string
		reqID int
		req   *entities.Book
		resp  entities.Book
		err   error
		err1  error
	}{
		{desc: "updating details", reqID: 1,
			req: &entities.Book{BookID: 1, Title: "it",
				Author:      entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "max"},
				Publication: "arihant", PublishedDate: "2020/10/23"},
			resp: entities.Book{BookID: 1, Title: "it",
				Author:      entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "max"},
				Publication: "arihant", PublishedDate: "2020/10/23"},
			err:  nil,
			err1: nil,
		},
	}
	for _, v := range testcases {
		mockCtrl := gomock.NewController(t)

		mockBookStore := datastore.NewMockBook(mockCtrl)
		mockAuthorStore := datastore.NewMockAuthor(mockCtrl)
		mock := New(mockBookStore, mockAuthorStore)

		mockAuthorStore.EXPECT().IncludeAuthor(context.TODO(), v.req.Author.AuthorID).Return(v.req.Author, v.err1).AnyTimes()
		mockBookStore.EXPECT().Putbook(context.TODO(), v.reqID, v.req).Return(v.resp, v.err).AnyTimes()

		res, err := mock.Putbook(context.TODO(), v.reqID, v.req)

		assert.Equal(t, v.err, err)
		assert.Equal(t, v.resp, res)
	}
}

func TestDeletebook(t *testing.T) {
	testcases := []struct {
		desc string
		req  int
		resp int64 // rowsAffected
		err  error
	}{
		{"success:valid id", 1, 1, nil},
		{"failure:invalid id", -4, 0, errors.New("invalid bookID")},
	}
	for _, v := range testcases {
		mockCtrl := gomock.NewController(t)

		mockAuthorStore := datastore.NewMockAuthor(mockCtrl)
		mockBookStore := datastore.NewMockBook(mockCtrl)
		mock := New(mockBookStore, mockAuthorStore)

		mockBookStore.EXPECT().Deletebook(context.TODO(), v.req).Return(v.resp, v.err).AnyTimes()
		res, err := mock.Deletebook(context.TODO(), v.req)

		assert.Equal(t, v.err, err)
		assert.Equal(t, v.resp, res)
	}
}
