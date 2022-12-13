package author

import (
	"context"
	"errors"
	"log"
	"testing"

	"GoLang-Interns-2022/entities"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestPostAuthor(t *testing.T) {
	testcases := []struct {
		desc         string
		req          entities.Author
		response     entities.Author
		lastInsertID int64
		rowsAffected int64
		err          error
	}{
		{"details of author with id 1",
			entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "max"},
			entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "max"},
			1,
			1,
			nil,
		},
		{"details of author with id 2",
			entities.Author{AuthorID: 2, Firstname: " mukhesh", Lastname: "mekala", Dob: "2000/10/24", Penname: "zero"},
			entities.Author{AuthorID: 2, Firstname: " mukhesh", Lastname: "mekala", Dob: "2000/10/24", Penname: "zero"},
			2,
			1,
			nil,
		},
		{"failure: invalid id ",
			entities.Author{AuthorID: -2, Firstname: " mukhesh", Lastname: "mekala", Dob: "2000/10/24", Penname: "zero"},
			entities.Author{},
			0,
			0,
			errors.New("invalid details"),
		},
		{"failure: author with id 2 exists",
			entities.Author{AuthorID: 2, Firstname: " mukhesh", Lastname: "mekala", Dob: "2000/10/24", Penname: "zero"},
			entities.Author{},
			0,
			0,
			errors.New("invalid details"),
		},
	}
	for _, v := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Print(err)
		}

		query := "insert into Author(authorId,Firstname,Lastname,Dob,Penname)values(?,?,?,?,?)"
		mock.ExpectExec(query).WithArgs(v.req.AuthorID, v.req.Firstname, v.req.Lastname, v.req.Dob, v.req.Penname).
			WillReturnResult(sqlmock.NewResult(v.lastInsertID, v.rowsAffected)).WillReturnError(v.err)

		authorStore := New(db)

		resp, err := authorStore.Postauthor(context.TODO(), v.req)

		assert.Equal(t, v.err, err)
		assert.Equal(t, v.response, resp)
	}
}

func TestPutAuthor(t *testing.T) {
	testcases := []struct {
		desc         string
		reqID        int
		req          entities.Author
		resp         entities.Author
		rows         *sqlmock.Rows
		lastInsertID int64
		rowsAffected int64
		err          error
	}{
		{desc: "success:updating author", reqID: 1,
			req:  entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "rolex"},
			resp: entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "rolex"},
			rows: sqlmock.NewRows([]string{"authorId", "Firstname", "Lastname", "Dob", "Penname"}).
				AddRow(1, "chetan", "bhagat", "1980/04/23", "rolex"),
			lastInsertID: 1,
			rowsAffected: 1,
			err:          nil,
		},
		{desc: "failure:invalid id", reqID: 111,
			req:          entities.Author{AuthorID: 111, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "rolex"},
			resp:         entities.Author{},
			rows:         sqlmock.NewRows([]string{}),
			lastInsertID: 0,
			rowsAffected: 0,
			err:          errors.New("author id does not exists"),
		},
	}
	for _, v := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Print(err)
		}

		mock.ExpectQuery("select * from Author where authorId=?").
			WithArgs(v.reqID).WillReturnRows(v.rows).WillReturnError(v.err)

		mock.ExpectExec("UPDATE Author SET Firstname=?,Lastname=?,Dob=?,Penname=? WHERE authorId=?").
			WithArgs(v.req.Firstname, v.req.Lastname, v.req.Dob, v.req.Penname, v.reqID).
			WillReturnResult(sqlmock.NewResult(v.lastInsertID, v.rowsAffected)).WillReturnError(v.err)

		authorStore := New(db)

		resp, err := authorStore.Putauthor(context.TODO(), v.reqID, v.req)

		assert.Equal(t, v.err, err)
		assert.Equal(t, v.resp, resp)
	}
}

func TestDeleteAuthor(t *testing.T) {
	testcases := []struct {
		desc         string
		req          int // authorID
		rows         *sqlmock.Rows
		lastInsertID int64
		rowsAffected int64
		err          error
		errorType    error
	}{
		{"success:deleting author ",
			1,
			sqlmock.NewRows([]string{"authorId", "Firstname", "Lastname", "Dob", "Penname"}).
				AddRow(1, "chetan", "bhagat", "1980/04/23", "rolex"),
			0,
			1,
			nil,
			nil,
		},
		{"failure:id already deleted",
			1,
			sqlmock.NewRows([]string{}),
			0,
			0,
			nil,
			errors.New("id not found"),
		},
	}
	for _, v := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Print(err)
		}

		query := "select * from Author where authorId=?"
		mock.ExpectQuery(query).WithArgs(v.req).WillReturnRows(v.rows).WillReturnError(v.err)

		query = "DELETE FROM Author WHERE authorId =?"
		mock.ExpectExec(query).WithArgs(v.req).WillReturnResult(sqlmock.NewResult(v.lastInsertID, v.rowsAffected)).
			WillReturnError(v.err)

		authorStore := New(db)

		res, err := authorStore.Deleteauthor(context.TODO(), v.req)

		assert.Equal(t, v.errorType, err)
		assert.Equal(t, v.rowsAffected, res)
	}
}

func TestIncludeAuthor(t *testing.T) {
	testcases := []struct {
		desc string
		req  int
		resp entities.Author
		row  *sqlmock.Rows
		err  error
	}{
		{"fetching author",
			1,
			entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "rolex"},
			sqlmock.NewRows([]string{"authorId", "Firstname", "Lastname", "Dob", "Penname"}).AddRow(1, "chetan", "bhagat", "1980/04/23", "rolex"),
			nil,
		},
		{"invalid author",
			-1,
			entities.Author{},
			sqlmock.NewRows([]string{}),
			errors.New("invalid author"),
		},
	}
	for _, v := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Print(err)
		}

		query := "select * from Author where authorId=?"
		mock.ExpectQuery(query).WithArgs(v.req).WillReturnRows(v.row).WillReturnError(v.err)

		authorStore := New(db)

		res, err := authorStore.IncludeAuthor(context.TODO(), v.req)

		assert.Equal(t, v.resp, res)
		assert.Equal(t, v.err, err)
	}
}
