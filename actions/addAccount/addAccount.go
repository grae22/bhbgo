package addAccount

import (
	"bhbgo/domain/accounts"
	"bhbgo/ui/selectAccount"
	"bhbgo/ui/selectAccountType"
	"fmt"
	"strings"
)

func Run() {
	var accountType *accounts.Account = selectAccountType.Run()

	if accountType == nil {
		return
	}

	accountsArr := []*accounts.Account{accountType}

	var parentAccount *accounts.Account = selectAccount.Run(
		accountsArr,
		"Select a parent account")

	if parentAccount == nil {
		return
	}

	name := getNewAccountName()

	if len(name) == 0 {
		return
	}

	newAccount := new(accounts.Account)
	newAccount.Name = name
	newAccount.AccountType = accountType.AccountType

	parentAccount.Children = append(parentAccount.Children, newAccount)
}

func getNewAccountName() string {
	fmt.Printf("\n> New account name: ")

	var name string
	fmt.Scanln(&name)

	return strings.Trim(name, " ")
}
