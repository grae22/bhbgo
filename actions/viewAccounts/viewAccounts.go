package viewAccounts

import (
	"bhbgo/domain/accounts"
	"bhbgo/ui/viewAccounts"
)

func Run() {

	viewAccounts.Run(accounts.RootsArr[:])
}
