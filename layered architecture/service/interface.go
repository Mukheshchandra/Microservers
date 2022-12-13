package service

import (
	"GoLang-Interns-2022/entities"
	"context"
)

type AuthorService interface {
	Postauthor(ctx context.Context, author entities.Author) (entities.Author, error)
	Putauthor(ctx context.Context, id int, author entities.Author) (entities.Author, error)
	Deleteauthor(ctx context.Context, id int) (int64, error)
}
type BookService interface {
	GetAllBook(ctx context.Context, title, includeAuthor string) ([]entities.Book, error)
	GetbyID(ctx context.Context, id int) (entities.Book, error)
	Postbook(ctx context.Context, book *entities.Book) (entities.Book, error)
	Putbook(ctx context.Context, id int, book *entities.Book) (entities.Book, error)
	Deletebook(ctx context.Context, id int) (int64, error)
}
