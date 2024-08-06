package models

import (
	"fmt"
	"time"
)

type Transaction struct {
	tracId string
	custId string
	isbn   string
	amt    float32
	tdate  time.Time
}

func NewTransaction(tracId string, custId string, isbn string, amt float32, tdate time.Time) *Transaction {
	return &Transaction{
		tracId: tracId,
		custId: custId,
		isbn:   isbn,
		amt:    amt,
		tdate:  tdate,
	}
}

func (t *Transaction) DisplayTransaction() {
	fmt.Println("Transaction Id : ", t.tracId)
	fmt.Println("Customer Id : ", t.custId)
	fmt.Println("ISBN : ", t.isbn)
	fmt.Println("Amount : ", t.amt)
	fmt.Println("Time : ", t.tdate)
}
