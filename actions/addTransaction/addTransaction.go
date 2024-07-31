package addTransaction

import (
	"bhbgo/domain/accounts"
	"bhbgo/ui/selectAccount"
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
}
