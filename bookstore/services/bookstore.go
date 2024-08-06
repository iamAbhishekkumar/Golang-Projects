package services

import (
	"bookstore/models"
	"math/rand"
	"strconv"
	"time"
)

type BookStore struct {
	inv          *Inventory
	Users        []*models.User
	transactions []*models.Transaction
	currUser     *models.User
}

func NewBookStore(inv *Inventory, users *models.User) *BookStore {
	return &BookStore{
		inv:   inv,
		Users: []*models.User{},
	}
}

func (bs *BookStore) AddUser(user *models.User) {
	bs.Users = append(bs.Users, user)
}

func (bs *BookStore) ShowInventory() error {
	return bs.inv.DisplayAllBooks()
}

func (bs *BookStore) AddTransaction(transaction *models.Transaction) {
	bs.transactions = append(bs.transactions, transaction)
}

func (bs *BookStore) BuyBook(qty int, isbn string) {
	if bs.inv.SearchAndPurchasable(qty, isbn) {
		bs.transactions = append(bs.transactions,
			models.NewTransaction(strconv.Itoa(int(rand.Int63())), bs.currUser.UserId(), isbn, rand.Float32(), time.Now()))
	}

}
