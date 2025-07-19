package main

import "math/rand"

type Account struct {
	ID        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Number    int64   `json:"number"`
	Email     string  `json:"email"`
	Balance   float64 `json:"balance"`
}

func NewAccount(firstName, lastName, email string) *Account {
	return &Account{
		ID:        rand.Intn(1000000), // Random ID for simplicity
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Number:    int64(rand.Intn(10000000)),
		Balance:   0,
	}
}
