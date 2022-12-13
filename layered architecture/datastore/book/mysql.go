package book

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"GoLang-Interns-2022/entities"
)

type Storer struct {
	db *sql.DB
}

func New(db *sql.DB) Storer {
	return Storer{db: db}
}

func (b Storer) GetAllBook(ctx context.Context) ([]entities.Book, error) {
	var books []entities.Book

	rows, err := b.db.QueryContext(ctx, "select * from Book")
	if err != nil {
		return []entities.Book{{}}, err
	}

	defer func() {
		_ = rows.Close()
		_ = rows.Err()
	}()

	for rows.Next() {
		var book entities.Book

		err := rows.Scan(&book.BookID, &book.Title, &book.Author.AuthorID, &book.Publication, &book.PublishedDate)
		if err != nil {
			log.Print(err)
		}

		books = append(books, book)
	}

	return books, nil
}

func (b Storer) GetBookbyTitle(ctx context.Context, title string) ([]entities.Book, error) {
	var books []entities.Book

	row, err := b.db.QueryContext(ctx, "select * from Book where Title=?", title)
	if err != nil {
		return []entities.Book{}, err
	}

	defer func() {
		_ = row.Close()
		_ = row.Err()
	}()

	for row.Next() {
		var book entities.Book
		if err := row.Scan(&book.BookID, &book.Title, &book.Author.AuthorID, &book.Publication, &book.PublishedDate); err != nil {
			log.Print(err)
		}

		books = append(books, book)
	}

	return books, nil
}

func (b Storer) GetbyID(ctx context.Context, id int) (entities.Book, error) {
	book := entities.Book{}

	bookrow := b.db.QueryRowContext(ctx, "SELECT * FROM Book WHERE bookId = ?", id)

	err := bookrow.Scan(&book.BookID, &book.Title, &book.Author.AuthorID, &book.Publication, &book.PublishedDate)
	if err != nil {
		return entities.Book{}, errors.New("invalid id")
	}

	return book, nil
}

func (b Storer) Postbook(ctx context.Context, book *entities.Book) (entities.Book, error) {
	res, err := b.db.ExecContext(ctx, "INSERT INTO Book(bookId,Title,authorId,Publication,PublishedDate) VALUES(?,?,?,?,?)",
		&book.BookID, &book.Title, &book.Author.AuthorID, &book.Publication, &book.PublishedDate)
	if err != nil {
		return entities.Book{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return entities.Book{}, err
	}

	book.BookID = int(id)

	return *book, nil
}

func (b Storer) Putbook(ctx context.Context, id int, book *entities.Book) (entities.Book, error) {
	row, err := b.db.QueryContext(ctx, "select * from Book where bookId=?", id)
	if err != nil {
		return entities.Book{}, err
	}

	defer func() {
		_ = row.Close()
		_ = row.Err()
	}()

	if !row.Next() {
		return entities.Book{}, errors.New("id does not exists")
	}

	_, err = b.db.ExecContext(ctx, "update  Book SET Title=?,authorId=?,Publication=?,PublishedDate=? where bookId=?",
		&book.Title, &book.Author.AuthorID, &book.Publication, &book.PublishedDate, id)
	if err != nil {
		return entities.Book{}, err
	}

	return *book, nil
}

func (b Storer) Deletebook(ctx context.Context, id int) (int64, error) {
	row, err := b.db.QueryContext(ctx, "select * from Book where bookId=?", id)
	if err != nil {
		return 0, err
	}

	defer func() {
		_ = row.Close()
		_ = row.Err()
	}()

	if !row.Next() {
		return 0, errors.New("id not found")
	}

	res, err := b.db.ExecContext(ctx, "DELETE FROM Book WHERE bookId =?", id)
	if err != nil {
		return 0, errors.New("invalid")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, errors.New("invalid")
	}

	return rowsAffected, nil
}
