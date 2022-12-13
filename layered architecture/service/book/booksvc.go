package book

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"GoLang-Interns-2022/datastore"
	"GoLang-Interns-2022/entities"
)

type Service struct {
	bookdatastore   datastore.Book
	authordatastore datastore.Author
}

func New(bd datastore.Book, ad datastore.Author) Service {
	return Service{bookdatastore: bd, authordatastore: ad}
}

func (b Service) GetAllBook(ctx context.Context, title, includeAuthor string) ([]entities.Book, error) {
	var (
		books []entities.Book
		err   error
	)

	if title != "" {
		books, err = b.bookdatastore.GetBookbyTitle(ctx, title)
		if err != nil {
			return []entities.Book{}, err
		}
	} else {
		books, err = b.bookdatastore.GetAllBook(ctx)
		if err != nil {
			return []entities.Book{}, err
		}
	}

	if includeAuthor == "true" {
		for i := range books {
			author, err := b.authordatastore.IncludeAuthor(ctx, books[i].Author.AuthorID)
			if err != nil {
				return []entities.Book{}, err
			}

			books[i].Author = author
		}
	}

	return books, nil
}

func (b Service) GetbyID(ctx context.Context, id int) (entities.Book, error) {
	if id <= 0 {
		return entities.Book{}, errors.New("invalid id")
	}

	book, err := b.bookdatastore.GetbyID(ctx, id)
	if err != nil {
		return entities.Book{}, errors.New("error")
	}

	return book, nil
}

func (b Service) Postbook(ctx context.Context, book *entities.Book) (entities.Book, error) {
	if book.BookID <= 0 || book.Title == "" || book.Author.AuthorID <= 0 || !publishedDate(book.PublishedDate) ||
		!publications(book.Publication) {
		return entities.Book{}, errors.New("invalid credentials")
	}

	existingAuthor, err := b.authordatastore.IncludeAuthor(ctx, book.Author.AuthorID)
	if err != nil {
		return entities.Book{}, err
	}

	_, err = b.bookdatastore.Postbook(ctx, book)
	if err != nil {
		return entities.Book{}, errors.New("error")
	}

	book.Author = existingAuthor

	return *book, nil
}

func (b Service) Putbook(ctx context.Context, id int, book *entities.Book) (entities.Book, error) {
	if id <= 0 || book.Title == "" || book.Author.AuthorID <= 0 || !publishedDate(book.PublishedDate) ||
		!publications(book.Publication) || book.Author.Firstname == "" || book.Author.Lastname == "" || book.Author.Penname == "" ||
		!checkDob(book.Author.Dob) {
		return entities.Book{}, errors.New("invalid credentials")
	}

	_, err := b.bookdatastore.Putbook(ctx, id, book)
	if err != nil {
		return entities.Book{}, errors.New("error")
	}

	return *book, nil
}

func (b Service) Deletebook(ctx context.Context, id int) (int64, error) {
	if id <= 0 {
		return 0, errors.New("invalid bookID")
	}

	count, err := b.bookdatastore.Deletebook(ctx, id)
	if err != nil {
		return 0, errors.New("error")
	}

	return count, nil
}

func publishedDate(publishedDate string) bool {
	date := strings.Split(publishedDate, "/")
	year, _ := strconv.Atoi(date[0])
	month, _ := strconv.Atoi(date[1])
	day, _ := strconv.Atoi(date[2])

	switch {
	case year <= 1880 || year > 2022:
		return false
	case month <= 0 || month > 12:
		return false
	case day <= 0 || day > 31:
		return false
	}

	return true
}

func publications(publications string) bool {
	switch strings.ToLower(publications) {
	case "penguin":
		return true
	case "scholastic":
		return true
	case "arihant":
		return true
	default:
		return false
	}
}

func checkDob(dOB string) bool {
	if dOB == "" {
		return false
	}

	dob := strings.Split(dOB, "/")

	year, _ := strconv.Atoi(dob[0])

	month, _ := strconv.Atoi(dob[1])

	day, _ := strconv.Atoi(dob[2])

	y := 2022

	switch {
	case year > y || year < 1000:
		return false
	case month <= 0 || month > 12:
		return false
	case day <= 0 || day > 31:
		return false
	}

	return true
}
