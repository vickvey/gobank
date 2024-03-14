package main

import (
	"math/big"
)

type Account struct {
	ID        big.Int
	FirstName string
	LastName  string
	Number    int64
	Balance   int64
}

/// TODO: correct the error here

// func NewAccount(firstName, lastName string) {
// 	return &Account{
// 		ID:        rand.Intn(10000),
// 		FirstName: firstName,
// 		LastName:  lastName,
// 		Number:    int64(rand.Intn(1000000)),
// 	}
// }
