package data

import (
	"database/sql"
	"fmt"
)

type Book struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookModel struct {
	DB *sql.DB
}

func (bookModel *BookModel) GetAll() ([]*Book, error) {
	taleOfTwoCities := &Book{
		ID:     1001,
		Title:  "A Tale of Two Cities",
		Author: "Charles Dickens",
	}

	theSelfishGene := &Book{
		ID:     1002,
		Title:  "The Selfish Gene",
		Author: "Richard Dawkins",
	}

	loveInTheTimeOfCholera := &Book{
		ID:     1003,
		Title:  "Love in the Time of Cholera",
		Author: "Gabriel Garcia Marquez",
	}

	return []*Book{taleOfTwoCities, theSelfishGene, loveInTheTimeOfCholera}, nil
}

func (bookModel *BookModel) GetById(id int64) (*Book, error) {
	books, err := bookModel.GetAll()
	if err != nil {
		return nil, err
	}

	for _, book := range books {
		if book.ID == id {
			return book, nil
		}
	}

	return nil, fmt.Errorf("failed to find book with id %d", id)
}
