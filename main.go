package main

import (
	"fmt"

	"github.com/usrname/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Luke")
	account.Deposit(10)
	fmt.Println(account)
}
