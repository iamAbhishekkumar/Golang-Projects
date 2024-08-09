package models

import "fmt"

type User struct {
	custName string
	userId   string
	email    string
}

func NewUser(custName string, userId string, email string) *User {
	return &User{
		custName: custName,
		userId:   userId,
		email:    email,
	}
}

func (u *User) UserId() string {
	return u.userId
}

func (u *User) Display() {
	fmt.Println("Customer Name : ", u.custName)
	fmt.Println("User Id : ", u.userId)
	fmt.Println("Email : ", u.email)

}
