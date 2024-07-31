package viewAccounts

import (
	"bhbgo/domain/accounts"
	"bhbgo/ui/utils/menu"
	"strconv"
)

func Run(rootAccounts []*accounts.Account) {

	var menuOptions []menu.MenuOption

	for _, account := range rootAccounts {
		buildAccountsMenuOptionsRecursive(
			account,
			0,
			&menuOptions)
	}

	wasCancelled := false

	menu.Show(
		"Accounts",
		menuOptions,
		"Select an account",
		"Back",
		&wasCancelled)
}

func buildAccountsMenuOptionsRecursive(
	account *accounts.Account,
	level int,
	menuOptions *[]menu.MenuOption) {

	newOption := menu.MenuOption{
		Text:   formatAccountName(account.Name, level) + " ... " + formatBalance(account),
		Action: nil,
		Value:  account,
	}

	*menuOptions = append(*menuOptions, newOption)

	if len(account.Children) == 0 {
		return
	}

	for _, child := range account.Children {
		buildAccountsMenuOptionsRecursive(
			child,
			level+1,
			menuOptions)
	}
}

func formatAccountName(
	name string,
	level int) string {

	if level == 0 {
		return name
	}

	prefix := "|-"

	for i := 1; i < level; i++ {
		prefix += "--"
	}

	return prefix + " " + name
}

func formatBalance(account *accounts.Account) string {

	balanceAbs := account.Balance

	if balanceAbs < 0 {
		balanceAbs = -balanceAbs
	}

	balanceStr := strconv.Itoa(balanceAbs)

	if account.Balance < 0 {
		return "[" + balanceStr + "]"
	}

	return balanceStr
}
