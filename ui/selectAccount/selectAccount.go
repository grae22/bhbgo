package selectAccount

import (
	"bhbgo/domain/accounts"
	"bhbgo/ui/utils/menu"
)

func Run(
	rootAccounts []*accounts.Account,
	menuInstruction string) *accounts.Account {

	menuOptions := []menu.MenuOption{}

	for i := range rootAccounts {
		buildMenuOptionsFromRootAccountRecursive(&menuOptions, rootAccounts[i], 0)
	}

	var wasCancelled bool

	selectedOption :=
		menu.Show(
			"Accounts",
			menuOptions,
			menuInstruction,
			"Cancel",
			&wasCancelled)

	if wasCancelled {
		return nil
	}

	return selectedOption.Value.(*accounts.Account)
}

func buildMenuOptionsFromRootAccountRecursive(
	menuOptions *[]menu.MenuOption,
	rootAccount *accounts.Account,
	level int) {

	if level == 0 {
		option := menu.MenuOption{
			Text:   formatMenuTextForHierarchy(rootAccount.Name, level),
			Action: nil,
			Value:  rootAccount,
		}

		*menuOptions = append(*menuOptions, option)
	}

	for _, account := range rootAccount.Children {
		option := menu.MenuOption{
			Text:   formatMenuTextForHierarchy(account.Name, level+1),
			Action: nil,
			Value:  account,
		}

		*menuOptions = append(*menuOptions, option)

		buildMenuOptionsFromRootAccountRecursive(
			menuOptions,
			account,
			level+1,
		)
	}
}

func formatMenuTextForHierarchy(
	text string,
	level int) string {

	if level == 0 {
		return text
	}

	var prefix = "|-"

	for i := 1; i < level; i++ {
		prefix += "--"
	}

	return prefix + " " + text
}
