package addTransaction

import (
	"bhbgo/console"
	"bhbgo/domain/accounts"
	"bhbgo/ui/selectAccount"
	"fmt"
	"strconv"
)

func Run() {

	var debitAccount *accounts.Account = selectAccount.Run(
		accounts.RootsArr[:],
		"Select an account to debit")

	if debitAccount == nil {
		return
	}

	var creditAccount *accounts.Account = selectAccount.Run(
		accounts.RootsArr[:],
		"Select an account to credit")

	if creditAccount == nil {
		return
	}

	console.Clear()

	fmt.Print("Enter an amount: ")

	var input string
	fmt.Scanln(&input)

	inputInt, _ := strconv.Atoi(input)

	updateBalance(debitAccount, -inputInt)
	updateBalance(creditAccount, inputInt)
}

func updateBalance(
	account *accounts.Account,
	amount int) {

	if account == nil {
		return
	}

	account.Balance += amount

	updateBalance(account.Parent, amount)
}
