package book

import (
	"context"
	"errors"
	"log"
	"testing"

	"GoLang-Interns-2022/entities"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetallbook(t *testing.T) {
	testcases := []struct {
		desc string
		resp []entities.Book
		rows *sqlmock.Rows
		err  error
	}{
		{desc: "get all books",
			resp: []entities.Book{
				{BookID: 1, Title: "it", Author: entities.Author{AuthorID: 1}, Publication: "arihant", PublishedDate: "2020/10/23"},
				{BookID: 2, Title: "god", Author: entities.Author{AuthorID: 2}, Publication: "arihant", PublishedDate: "2015/07/03"}},
			rows: sqlmock.NewRows([]string{"bookId", "Title", "authorId", "Publication", "PublishedDate"}).
				AddRow(1, "it", 1, "arihant", "2020/10/23").
				AddRow(2, "god", 2, "arihant", "2015/07/03"),
			err: nil,
		},
	}
	for _, v := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Print(err)
		}

		query := "select * from Book"
		mock.ExpectQuery(query).WillReturnRows(v.rows).WillReturnError(v.err)

		bookStore := New(db)

		resp, err := bookStore.GetAllBook(context.TODO())
		assert.Equal(t, v.resp, resp)
		assert.Equal(t, v.err, err)
	}
}

func TestGetBookByTitle(t *testing.T) {
	testcases := []struct {
		desc string
		req  string
		resp []entities.Book
		row  *sqlmock.Rows
		err  error
	}{
		{desc: "fetching details of book by title",
			req: "it",
			resp: []entities.Book{{BookID: 1, Title: "it", Author: entities.Author{AuthorID: 1},
				Publication: "arihant", PublishedDate: "2020/10/23"}},
			row: sqlmock.NewRows([]string{"bookId", "Title", "authorId", "Publication", "PublishedDate"}).
				AddRow(1, "it", 1, "arihant", "2020/10/23"),
			err: nil,
		},
	}
	for _, v := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Print(err)
		}

		query := "select * from Book where Title=?"
		mock.ExpectQuery(query).WithArgs(v.req).WillReturnRows(v.row).WillReturnError(v.err)

		bookStore := New(db)

		resp, err := bookStore.GetBookbyTitle(context.TODO(), v.req)

		assert.Equal(t, v.err, err)
		assert.Equal(t, v.resp, resp)
	}
}

func TestGetbyID(t *testing.T) {
	testcases := []struct {
		desc           string
		req            int
		resp           entities.Book
		row            *sqlmock.Rows
		lastInsertedID int64
		rowsAffected   int64
		err            error
	}{
		{desc: "fetching details of book by id",
			req: 1,
			resp: entities.Book{BookID: 1, Title: "it", Author: entities.Author{AuthorID: 1},
				Publication: "arihant", PublishedDate: "2020/10/23"},
			row: sqlmock.NewRows([]string{"bookId", "Title", "authorId", "Publication", "PublishedDate"}).
				AddRow(1, "it", 1, "arihant", "2020/10/23"),
			lastInsertedID: 0,
			rowsAffected:   0,
			err:            nil,
		},
		{"invaid id",
			-1,
			entities.Book{},
			sqlmock.NewRows([]string{}),
			0,
			0,
			errors.New("invalid id"),
		},
	}
	for _, v := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Print(err)
		}

		query := "SELECT * FROM Book WHERE bookId = ?"
		mock.ExpectQuery(query).WithArgs(v.req).WillReturnRows(v.row).WillReturnError(v.err)

		bookStore := New(db)

		resp, err := bookStore.GetbyID(context.TODO(), v.req)

		assert.Equal(t, v.err, err)
		assert.Equal(t, v.resp, resp)
	}
}

func TestPostbook(t *testing.T) {
	testcases := []struct {
		desc           string
		req            entities.Book
		resp           entities.Book
		lastInsertedID int64
		rowsAffected   int64
		err            error
	}{
		{desc: "valid book",
			req: entities.Book{BookID: 1, Title: "2states", Author: entities.Author{AuthorID: 1,
				Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "rolex"}, Publication: "arihant",
				PublishedDate: "2020/10/23"},
			resp: entities.Book{BookID: 1, Title: "2states", Author: entities.Author{AuthorID: 1, Firstname: "chetan",
				Lastname: "bhagat", Dob: "1980/04/23", Penname: "rolex"}, Publication: "arihant", PublishedDate: "2020/10/23"},
			lastInsertedID: 1,
			rowsAffected:   1,
			err:            nil,
		},
		{
			desc: "valid book",
			req: entities.Book{BookID: 2, Title: "god", Author: entities.Author{AuthorID: 2, Firstname: "mukhesh", Lastname: "mekala",
				Dob: "2000/10/24", Penname: "zero"}, Publication: "arihant", PublishedDate: "2015/07/03"},
			resp: entities.Book{BookID: 2, Title: "god", Author: entities.Author{AuthorID: 2, Firstname: "mukhesh", Lastname: "mekala",
				Dob: "2000/10/24", Penname: "zero"}, Publication: "arihant", PublishedDate: "2015/07/03"},
			lastInsertedID: 2,
			rowsAffected:   1,
			err:            nil,
		},
	}
	for _, v := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Print(err)
		}

		query := "INSERT INTO Book(bookId,Title,authorId,Publication,PublishedDate) VALUES(?,?,?,?,?)"
		mock.ExpectExec(query).WithArgs(v.req.BookID, v.req.Title, v.req.Author.AuthorID, v.req.Publication, v.req.PublishedDate).
			WillReturnResult(sqlmock.NewResult(v.lastInsertedID, v.rowsAffected)).WillReturnError(v.err)

		bookStore := New(db)

		resp, err := bookStore.Postbook(context.TODO(), &v.req)
		assert.Equal(t, v.err, err)
		assert.Equal(t, v.resp, resp)
	}
}

func TestPutbook(t *testing.T) {
	testcases := []struct {
		desc           string
		reqID          int
		req            entities.Book
		resp           entities.Book
		rows           *sqlmock.Rows
		lastInsertedID int64
		rowsAffected   int64
		err            error
	}{
		{desc: "updating details", reqID: 1,
			req: entities.Book{BookID: 1, Title: "it",
				Author:      entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "max"},
				Publication: "arihant", PublishedDate: "2020/10/23"},
			resp: entities.Book{BookID: 1, Title: "it",
				Author:      entities.Author{AuthorID: 1, Firstname: "chetan", Lastname: "bhagat", Dob: "1980/04/23", Penname: "max"},
				Publication: "arihant", PublishedDate: "2020/10/23"},
			rows: sqlmock.NewRows([]string{"bookId", "Title", "authorId", "Publication", "PublishedDate"}).
				AddRow(1, "it", 1, "arihant", "2020/10/23"),
			lastInsertedID: 1,
			rowsAffected:   1,
			err:            nil,
		},
	}
	for _, v := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Print(err)
		}

		query := "select * from Book where bookId=?"
		mock.ExpectQuery(query).WithArgs(v.reqID).WillReturnRows(v.rows).WillReturnError(v.err)

		query = "update  Book SET Title=?,authorId=?,Publication=?,PublishedDate=? where bookId=?"
		mock.ExpectExec(query).WithArgs(v.req.Title, v.req.Author.AuthorID, v.req.Publication, v.req.PublishedDate, v.reqID).
			WillReturnResult(sqlmock.NewResult(v.lastInsertedID, v.rowsAffected)).WillReturnError(v.err)

		bookStore := New(db)

		resp, err := bookStore.Putbook(context.TODO(), v.reqID, &v.req)

		assert.Equal(t, v.err, err)
		assert.Equal(t, v.resp, resp)
	}
}

func TestDeletebook(t *testing.T) {
	testcases := []struct {
		desc           string
		req            int
		rows           *sqlmock.Rows
		lastInsertedID int64
		rowsAffected   int64
		err            error
		errID          error
	}{
		{"success:valid id",
			1,
			sqlmock.NewRows([]string{"bookId", "Title", "authorId", "Publication", "PublishedDate"}).
				AddRow(1, "it", 1, "arihant", "2020/10/23"),
			0,
			1,
			nil,
			nil,
		},
		{"failure:id does not exists",
			4,
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

		query := "select * from Book where bookId=?"
		mock.ExpectQuery(query).WithArgs(v.req).WillReturnRows(v.rows).WillReturnError(v.err)

		query = "DELETE FROM Book WHERE bookId =?"
		mock.ExpectExec(query).WithArgs(v.req).WillReturnResult(sqlmock.NewResult(v.lastInsertedID, v.rowsAffected)).
			WillReturnError(v.err)

		bookStore := New(db)

		res, err := bookStore.Deletebook(context.TODO(), v.req)

		assert.Equal(t, v.rowsAffected, res)
		assert.Equal(t, v.errID, err)
	}
}
