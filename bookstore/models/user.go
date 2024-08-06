package models

type User struct {
	custName string
	custId   string
	email    string
}

func NewUser(custName string, custId string, email string) *User {
	return &User{
		custName: custName,
		custId:   custId,
		email:    email,
	}
}
