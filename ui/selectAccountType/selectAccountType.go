package selectAccountType

import (
	"bhbgo/domain/accounts"
	"bhbgo/ui/utils/menu"
)

func Run() *accounts.Account {
	var wasCancelled bool

	selectedOption := menu.Show(
		"Account Types",
		buildMenuOptionsFromAccountTypes(),
		"Select an account type",
		"Cancel",
		&wasCancelled)

	if wasCancelled {
		return nil
	}

	return selectedOption.Value.(*accounts.Account)
}

func buildMenuOptionsFromAccountTypes() []menu.MenuOption {
	var menuOptions []menu.MenuOption

	for i, accountType := range accounts.AccountTypes {
		option := menu.MenuOption{
			Text:   accountType,
			Action: nil,
			Value:  accounts.RootsArr[i],
		}

		menuOptions = append(menuOptions, option)
	}

	return menuOptions
}
