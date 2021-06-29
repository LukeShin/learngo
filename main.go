package main

import (
	"fmt"

	"github.com/learngo/banking"
)

func main() {
	account := banking.Account{Owner: "Luke", Balance: 10000000000}
	fmt.Println(account)
}
