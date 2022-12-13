package datastore

import (
	"GoLang-Interns-2022/entities"
	"context"
)

type Author interface {
	Postauthor(ctx context.Context, author entities.Author) (entities.Author, error)
	Putauthor(ctx context.Context, id int, author entities.Author) (entities.Author, error)
	Deleteauthor(ctx context.Context, id int) (int64, error)
	IncludeAuthor(ctx context.Context, id int) (entities.Author, error)
}

type Book interface {
	GetAllBook(ctx context.Context) ([]entities.Book, error)
	GetBookbyTitle(ctx context.Context, title string) ([]entities.Book, error)
	GetbyID(ctx context.Context, id int) (entities.Book, error)
	Postbook(ctx context.Context, book *entities.Book) (entities.Book, error)
	Putbook(ctx context.Context, id int, book *entities.Book) (entities.Book, error)
	Deletebook(ctx context.Context, id int) (int64, error)
}
