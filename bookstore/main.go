package main

import (
	"bookstore/services"
	"fmt"
)

func main() {
	inv := services.NewEmptyInventory()
	bookStore := services.NewBookStore(inv, nil)
	err := bookStore.ShowInventory()
	if err != nil {
		fmt.Println(err.Error())
	}
}
