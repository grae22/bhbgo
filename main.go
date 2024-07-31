package main

import (
	"bhbgo/console"
	"bhbgo/domain/accounts"
	"bhbgo/ui/mainMenu"
	"fmt"
)

func main() {
	accounts.Initialise()

	mainMenu.Run()

	console.Clear()

	fmt.Print("Bye.\n\n")
}
