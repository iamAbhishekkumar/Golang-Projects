package services

import (
	"bookstore/helpers"
	"bookstore/models"
)

type Inventory struct {
	books []models.Book
}

func NewEmptyInventory() *Inventory {
	return &Inventory{}
}

func (b *Inventory) AddBook(title string, author string, isbn string, price float32, qty int) error {
	b.books = append(b.books, *models.NewBook(title, author, isbn, price, qty))
	return nil
}

func (b *Inventory) SearchBookByISBN(isbn string) (int, error) {
	for idx, book := range b.books {
		if book.GetIsbn() == isbn {
			return idx, nil
		}
	}
	return -1, &helpers.BookNotFoundError{}
}

func (b *Inventory) SearchBookByTitle(title string) (int, error) {
	for idx, book := range b.books {
		if book.GetTitle() == title {
			return idx, nil
		}
	}
	return -1, &helpers.BookNotFoundError{}
}

func (b *Inventory) DeleteBook(isbn string) error {
	idx, err := b.SearchBookByISBN(isbn)
	if err != nil {
		return err
	}
	b.books = append(b.books[:idx], b.books[idx+1:]...)
	return nil
}

func (b *Inventory) UpdateBookQty(isbn string, qty int) error {
	idx, err := b.SearchBookByISBN(isbn)
	if err != nil {
		return nil
	}
	b.books[idx].SetQty(qty)
	return nil
}

func (b *Inventory) DisplayAllBooks() error {
	if len(b.books) == 0 {
		return &helpers.NoBooksInInventory{}
	}
	for _, book := range b.books {
		book.DisplayBook()
	}
	return nil
}
