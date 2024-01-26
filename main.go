package main

import (
	"fmt"
)

func printMenu() {
	fmt.Println("Please select an option:")

	for i := 1; i <= len(menu); i++ {
		fmt.Printf("%d. %s\n", i, menu[i].name)
	}
}

func main() {
	initMenu()
	printMenu()

	initMissles()

	for {
		var choice int

		fmt.Scanln(&choice)

		if item, exists := menu[choice]; exists {
			item.action()
		} else {
			// Clear the screen
			fmt.Print("\033[H\033[2J")
			printMenu()
		}
	}
}
