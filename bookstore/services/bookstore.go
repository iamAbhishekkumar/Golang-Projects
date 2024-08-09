package services

import (
	"bookstore/helpers"
	"bookstore/models"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type BookStore struct {
	inv          *Inventory
	users        []*models.User
	transactions []*models.Transaction
	currUser     *models.User
}

func NewBookStore(inv *Inventory, users *models.User) *BookStore {
	return &BookStore{
		inv:   inv,
		users: []*models.User{},
	}
}

func (bs *BookStore) AddUser(user *models.User) {
	bs.users = append(bs.users, user)
	*bs.currUser = *user
}

func (bs *BookStore) SearchUser(userId string) int {
	for idx, u := range bs.users {
		if u.UserId() == userId {
			return idx
		}
	}
	return -1
}

func (bs *BookStore) DeleteUser(userId string) {
	userToDelIdx := bs.SearchUser(userId)
	if userToDelIdx != -1 {
		bs.users = append(bs.users[:userToDelIdx], bs.users[userToDelIdx+1:]...)
	}
}

func (bs *BookStore) DisplayAllUsers() {
	if len(bs.users) == 0 {
		fmt.Println("No User Found")
		return
	}
	for _, u := range bs.users {
		u.Display()
		fmt.Println()
	}
}

func (bs *BookStore) ShowInventory() error {
	return bs.inv.DisplayAllBooks()
}

func (bs *BookStore) AddTransaction(transaction *models.Transaction) {
	bs.transactions = append(bs.transactions, transaction)
}

func (bs *BookStore) BuyBook(purchaseQty int, isbn string) error {
	idx, err := bs.inv.SearchBookByISBN(isbn)
	if err != nil {
		return err
	}
	bookOrigQty := bs.inv.books[idx].GetQty()
	if bookOrigQty >= purchaseQty {
		bs.transactions = append(bs.transactions,
			models.NewTransaction(strconv.Itoa(int(rand.Int63())), bs.currUser.UserId(), isbn, rand.Float32(), time.Now()))
		bookOrigQty -= purchaseQty
		bs.inv.UpdateBookQty(isbn, bookOrigQty)
		return nil
	}
	return &helpers.NotEnoughBooksQty{}
}
