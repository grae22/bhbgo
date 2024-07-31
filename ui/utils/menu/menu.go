package menu

import (
	"bhbgo/console"
	"fmt"
	"strconv"
	"strings"
)

type MenuOption struct {
	Text   string
	Action func()
	Value  interface{}
}

const CancelKey byte = '0'

type menu struct {
	title           string
	options         []MenuOption
	instructionText string
	cancelText      string
}

func Show(
	title string,
	options []MenuOption,
	instructionText string,
	cancelText string,
	wasCancelled *bool) *MenuOption {

	var menu = new(menu)
	menu.title = title
	menu.options = options
	menu.instructionText = instructionText
	menu.cancelText = cancelText

	drawMenu(menu)

	return handleInput(menu, wasCancelled)
}

func drawMenu(menu *menu) {
	clear()
	drawHeader(menu)
	drawOptions(menu)
	drawFooter(menu)
}

func clear() {
	console.Clear()
}

func drawHeader(menu *menu) {
	fmt.Printf("- %s -\n", menu.title)
}

func drawOptions(menu *menu) {
	for index, opt := range menu.options {
		fmt.Printf("\n%d : %s", index+1, opt.Text)
	}

	if len(menu.options) == 0 {
		fmt.Printf("\n(none)")
	}

	fmt.Printf("\n\n%c : %s", CancelKey, menu.cancelText)
}

func drawFooter(menu *menu) {
	fmt.Printf("\n\n> %s: ", menu.instructionText)
}

func handleInput(
	menu *menu,
	wasCancelled *bool) *MenuOption {

	for {
		var input string
		fmt.Scanln(&input)

		if len(input) > 0 &&
			strings.ToUpper(input)[0] == CancelKey {

			*wasCancelled = true
			break
		}

		var selectedIndex, err = strconv.Atoi(input)

		if err == nil {
			isValidSelection, selectedOption := executeSelectOption(menu, selectedIndex)

			if isValidSelection && selectedOption.Action == nil {
				return selectedOption
			}

			drawMenu(menu)
		}
	}

	return nil
}

func executeSelectOption(
	menu *menu,
	selectedIndex int) (bool, *MenuOption) {

	// Decrement since the displayed indices are 1-based and the
	// options array is (of course) zero-based.
	selectedIndex--

	if selectedIndex >= len(menu.options) {
		return false, nil
	}

	selectedOption := &menu.options[selectedIndex]

	if selectedOption.Action != nil {
		selectedOption.Action()
	}

	return true, selectedOption
}
