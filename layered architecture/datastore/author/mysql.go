package author

import (
	"context"
	"database/sql"
	"errors"

	"GoLang-Interns-2022/entities"
)

type Storer struct {
	db *sql.DB
}

func New(db *sql.DB) Storer {
	return Storer{db: db}
}

func (a Storer) Postauthor(ctx context.Context, author entities.Author) (entities.Author, error) {
	res, err := a.db.ExecContext(ctx, "insert into Author(authorId,Firstname,Lastname,Dob,Penname)values(?,?,?,?,?)",
		&author.AuthorID, &author.Firstname, &author.Lastname, &author.Dob, &author.Penname)
	if err != nil {
		return entities.Author{}, errors.New("invalid details")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return entities.Author{}, err
	}

	author.AuthorID = int(id)

	return author, nil
}

func (a Storer) Putauthor(ctx context.Context, id int, author entities.Author) (entities.Author, error) {
	row, err := a.db.QueryContext(ctx, "select * from Author where authorId=?", id)
	if err != nil {
		return entities.Author{}, err
	}

	defer func() {
		_ = row.Close()
		_ = row.Err()
	}()

	if !row.Next() {
		return entities.Author{}, errors.New("author id does not exists")
	}

	_, err = a.db.ExecContext(ctx, "UPDATE Author SET Firstname=?,Lastname=?,Dob=?,Penname=? WHERE authorId=?",
		&author.Firstname, &author.Lastname, &author.Dob, &author.Penname, id)
	if err != nil {
		return entities.Author{}, err
	}

	return author, nil
}

func (a Storer) Deleteauthor(ctx context.Context, id int) (int64, error) {
	row, err := a.db.QueryContext(ctx, "select * from Author where authorId=?", id)
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

	res, err := a.db.ExecContext(ctx, "DELETE FROM Author WHERE authorId =?", id)
	if err != nil {
		return 0, errors.New("invalid")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, errors.New("invalid")
	}

	return rowsAffected, nil
}

func (a Storer) IncludeAuthor(ctx context.Context, id int) (entities.Author, error) {
	row := a.db.QueryRowContext(ctx, "select * from Author where authorId=?", id)

	var author entities.Author

	if err := row.Scan(&author.AuthorID, &author.Firstname, &author.Lastname, &author.Dob, &author.Penname); err != nil {

		return entities.Author{}, errors.New("invalid author")
	}

	return author, nil
}
