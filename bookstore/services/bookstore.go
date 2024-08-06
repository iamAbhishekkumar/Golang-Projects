package services

import "bookstore/models"

type BookStore struct {
	inv   *Inventory
	users []*models.User
}

func NewBookStore(inv *Inventory, users *models.User) *BookStore {
	return &BookStore{
		inv:   inv,
		users: []*models.User{},
	}
}

func (bs *BookStore) AddUser(user *models.User) {
	bs.users = append(bs.users, user)
}

func (bs *BookStore) ShowInventory() error {
	return bs.inv.DisplayAllBooks()
}
