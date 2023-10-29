package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(1000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000)),
		CreatedAt: time.Now().UTC(),
	}
}
