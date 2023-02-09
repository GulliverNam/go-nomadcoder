package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates a new account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
// a is receiver and that name is the first lowercase letter of struct
// if you want balance of the original account, reciver is the pointer variable
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance returns the balance of the account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) WithDraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner changes the owner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner returns the owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
