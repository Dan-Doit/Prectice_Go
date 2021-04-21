package main

import (
	"fmt"
	"log"

	"github.com/Dan-Doit/prectice-go/bank"
)

func main() {
	account := bank.CreateAccount("Dan")
	account.Deposit(1000)

	err := account.Withdraw(100)

	// err print then exit
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account)
}
