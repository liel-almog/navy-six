package main

import (
	"fmt"
)

func main() {
	printMenu()
	initLaunchers()

	for {
		fmt.Println("Enter your choice: ")
		var choice int
		fmt.Scanln(&choice)

		if item, ok := menu[menuOption(choice)]; ok {
			item.action()
		} else {
			fmt.Println("Invalid choice")
		}

	}
}
