package author

import (
	"context"
	"errors"
	"testing"

	"GoLang-Interns-2022/datastore"
	"GoLang-Interns-2022/entities"

	"github.com/golang/mock/gomock"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/stretchr/testify/assert"
)

func TestPostauthor(t *testing.T) {
	testcases := []struct {
		desc string
		req  entities.Author
		resp entities.Author
		err  error
	}{
		{"valid author",
			entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "max"},
			entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "max"},
			nil,
		},
		{"failure:invalid authorID",
			entities.Author{AuthorID: -1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "max"},
			entities.Author{},
			errors.New("invalid authorID"),
		},
		{"failure:empty author first name",
			entities.Author{AuthorID: 2, Firstname: "", Lastname: "bhagat", Dob: "1980/04/23", Penname: "max"},
			entities.Author{},
			errors.New("invalid credentials"),
		},
		{"failure:empty author last name",
			entities.Author{AuthorID: 3, Firstname: "chetan", Lastname: "", Dob: "1980/04/23", Penname: "max"},
			entities.Author{},
			errors.New("invalid credentials"),
		},
		{"failure:empty author pen name",
			entities.Author{AuthorID: 4, Firstname: "chetan", Lastname: "king", Dob: "1980/04/23", Penname: ""},
			entities.Author{},
			errors.New("invalid credentials"),
		},
		{"failure:invalid DOB",
			entities.Author{AuthorID: 5, Firstname: "chetan", Lastname: "king", Dob: "19/04/23", Penname: "max"},
			entities.Author{},
			errors.New("invalid DOB"),
		},
		{"failure:invalid DOB",
			entities.Author{AuthorID: 5, Firstname: "chetan", Lastname: "king", Dob: "", Penname: "max"},
			entities.Author{},
			errors.New("invalid DOB"),
		},
		{"failure:invalid DOB",
			entities.Author{AuthorID: 5, Firstname: "chetan", Lastname: "king", Dob: "1900/0/23", Penname: "max"},
			entities.Author{},
			errors.New("invalid DOB"),
		},
		{"failure:invalid DOB",
			entities.Author{AuthorID: 5, Firstname: "chetan", Lastname: "king", Dob: "1900/04/0", Penname: "max"},
			entities.Author{},
			errors.New("invalid DOB"),
		},
	}
	for _, v := range testcases {
		mockCtrl := gomock.NewController(t)

		mockAuthorDatastore := datastore.NewMockAuthor(mockCtrl)
		mock := New(mockAuthorDatastore)

		mockAuthorDatastore.EXPECT().Postauthor(context.TODO(), v.req).Return(v.resp, v.err).AnyTimes()
		res, err := mock.Postauthor(context.TODO(), v.req)

		assert.Equal(t, v.resp, res)
		assert.Equal(t, v.err, err)
	}
}

func TestPutauthor(t *testing.T) {
	testcases := []struct {
		desc  string
		reqID int
		req   entities.Author
		resp  entities.Author
		err   error
	}{
		{desc: "valid author", reqID: 1,
			req:  entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "rolex"},
			resp: entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "rolex"},
			err:  nil,
		},
		{desc: "invalid authorID", reqID: -1,
			req:  entities.Author{AuthorID: -1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "rolex"},
			resp: entities.Author{},
			err:  errors.New("invalid authorID"),
		},
		{desc: "empty first name", reqID: 2,
			req:  entities.Author{AuthorID: 2, Firstname: "", Lastname: "bhagat", Dob: "1980/04/23", Penname: "rolex"},
			resp: entities.Author{},
			err:  errors.New("invalid credentials"),
		},
		{desc: "empty last name", reqID: 3,
			req:  entities.Author{AuthorID: 3, Firstname: "chic", Lastname: "", Dob: "1980/04/23", Penname: "rolex"},
			resp: entities.Author{},
			err:  errors.New("invalid credentials"),
		},
		{desc: "empty pen name", reqID: 4,
			req:  entities.Author{AuthorID: 4, Firstname: "mukhesh", Lastname: "bhagat", Dob: "1980/04/23", Penname: ""},
			resp: entities.Author{},
			err:  errors.New("invalid credentials"),
		},
		{desc: "invalid DOb", reqID: 5,
			req:  entities.Author{AuthorID: 5, Firstname: "chetan", Lastname: "bhagat", Dob: "19/04/23", Penname: "rolex"},
			resp: entities.Author{},
			err:  errors.New("invalid DOB"),
		},
	}
	for _, v := range testcases {
		mockCtrl := gomock.NewController(t)

		mockAuthorDataStore := datastore.NewMockAuthor(mockCtrl)
		mock := New(mockAuthorDataStore)

		mockAuthorDataStore.EXPECT().Putauthor(context.TODO(), v.reqID, v.req).Return(v.resp, v.err).AnyTimes()
		res, err := mock.Putauthor(context.TODO(), v.reqID, v.req)

		assert.Equal(t, v.resp, res)
		assert.Equal(t, v.err, err)
	}
}

func TestDeleteauthor(t *testing.T) {
	testcases := []struct {
		desc string
		req  int
		resp int64 // rowsAffected
		err  error
	}{
		{"valid id", 1, 1, nil},
		{"invalid id", -1, 0, errors.New("invalid id")},
		{"id not found", 161, 0, errors.New("id do not exists")},
	}
	for _, v := range testcases {
		mockCtrl := gomock.NewController(t)
		mockAuthorDataBase := datastore.NewMockAuthor(mockCtrl)
		mock := New(mockAuthorDataBase)

		mockAuthorDataBase.EXPECT().Deleteauthor(context.TODO(), v.req).Return(v.resp, v.err).AnyTimes()

		res, err := mock.Deleteauthor(context.TODO(), v.req)

		assert.Equal(t, v.resp, res)
		assert.Equal(t, v.err, err)
	}
}
