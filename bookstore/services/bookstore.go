package services

import "fmt"

type BookStore struct {
	books []Book
}

func NewBookStore() *BookStore {
	return &BookStore{}
}

func (b *BookStore) AddBook(title string, author string, isbn string, price float32, qty int) error {
	b.books = append(b.books, *NewBook(
		title, author, isbn, price, qty,
	))
	return nil
}
func (b *BookStore) SearchBook(isbn string) (int, error) {
	for idx, book := range b.books {
		if book.isbn == isbn {
			return idx, nil
		}
	}
	return -1, &BookNotFoundError{}
}

func (b *BookStore) DeleteBook(isbn string) error {
	idx, err := b.SearchBook(isbn)
	if err != nil {
		return err
	}
	b.books = append(b.books[:idx], b.books[idx+1:]...)
	return nil
}

func (b *BookStore) UpdateBook(isbn string) error {
	fmt.Print("Updtae BOoks")
	return nil
}

func (b *BookStore) DisplayAllBooks() error {
	if len(b.books) == 0 {
		return &NoBooksInInventory{}
	}
	for _, book := range b.books {
		book.DisplayBook()
	}
	return nil
}
