package services

import "fmt"

type Book struct {
	title  string
	author string
	isbn   string
	price  float32
	qty    int
}

func NewBook(title string, author string, isbn string, price float32, qty int) *Book {
	return &Book{
		title:  title,
		author: author,
		isbn:   isbn,
		price:  price,
		qty:    qty,
	}
}

func (b *Book) DisplayBook() {
	fmt.Println("Name : ", b.title)
	fmt.Println("Author : ", b.author)
	fmt.Println("ISBN : ", b.isbn)
	fmt.Println("Price : ", b.price)
	fmt.Println("Quantity : ", b.qty)
}
