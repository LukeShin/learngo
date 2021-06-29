package main

import (
	"fmt"

	"github.com/usrname/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Luke")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
}
