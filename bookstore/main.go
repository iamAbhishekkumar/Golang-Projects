package main

import (
	"bookstore/services"
	"fmt"
)

func main() {
	bookStore := services.NewBookStore()
	err := bookStore.DisplayAllBooks()
	if err != nil {
		fmt.Println(err.Error())
	}
}
