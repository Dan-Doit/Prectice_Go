package bank

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

// error when doesn't have money
var errNoMoney = errors.New("Sorry, Not enough money")

// CreateAccount is create user account
func CreateAccount(name string) *Account {
	account := Account{owner: name, balance: 0}
	return &account
}

// Deposit save amount
func (a *Account) Deposit(amount int) int {
	a.balance += amount
	return a.balance
}

// Balance check Balance
func (a Account) Balance() Account {
	return a
}

// Withdraw withdraw money
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	} else {
		a.balance -= amount
		return nil
	}
}

// to String by Default when called account
func (a Account) String() string {
	return fmt.Sprint(a.owner, "'s account\nhas : ", a.balance, "$")
}
