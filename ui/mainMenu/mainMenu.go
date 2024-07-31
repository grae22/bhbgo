package mainMenu

import (
	"bhbgo/actions/addAccount"
	"bhbgo/actions/addTransaction"
	"bhbgo/actions/viewAccounts"
	"bhbgo/ui/utils/menu"
)

func Run() {

	menuOptions := []menu.MenuOption{
		{Text: "Add account", Action: addAccount.Run},
		{Text: "View accounts", Action: viewAccounts.Run},
		{Text: "New transaction", Action: addTransaction.Run},
	}

	selectedQuit := false

	menu.Show(
		"Main Menu",
		menuOptions,
		"Enter your selection",
		"Quit",
		&selectedQuit)
}
