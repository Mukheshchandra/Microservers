package author

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"GoLang-Interns-2022/datastore"
	"GoLang-Interns-2022/entities"
)

type Authorservice struct {
	datastore datastore.Author
}

func New(d datastore.Author) Authorservice {
	return Authorservice{datastore: d}
}

func (a Authorservice) Postauthor(ctx context.Context, author entities.Author) (entities.Author, error) {
	if author.Firstname == "" || author.Lastname == "" || author.Penname == "" {
		return entities.Author{}, errors.New("invalid credentials")
	}

	if !checkDob(author.Dob) {
		return entities.Author{}, errors.New("invalid DOB")
	}

	if author.AuthorID <= 0 {
		return entities.Author{}, errors.New("invalid authorID")
	}

	_, _ = a.datastore.Postauthor(ctx, author)

	return author, nil
}

func (a Authorservice) Putauthor(ctx context.Context, id int, author entities.Author) (entities.Author, error) {
	if author.Firstname == "" || author.Lastname == "" || author.Penname == "" {
		return entities.Author{}, errors.New("invalid credentials")
	}

	if !checkDob(author.Dob) {
		return entities.Author{}, errors.New("invalid DOB")
	}

	if id <= 0 {
		return entities.Author{}, errors.New("invalid authorID")
	}

	_, _ = a.datastore.Putauthor(ctx, id, author)

	return author, nil
}

func (a Authorservice) Deleteauthor(ctx context.Context, id int) (int64, error) {
	if id <= 0 {
		return 0, errors.New("invalid id")
	}

	count, err := a.datastore.Deleteauthor(ctx, id)
	if err != nil {
		return 0, errors.New("id do not exists")
	}

	return count, nil
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
