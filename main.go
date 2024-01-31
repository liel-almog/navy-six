package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lielalmog/navy-six/menu"
	"github.com/lielalmog/navy-six/missiles"
)

func cleanScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	r := bufio.NewReader(os.Stdin)
	missiles.InitLaunchers()

	for {
		cleanScreen()
		menu.PrintMenu()

		fmt.Print("Enter your choice: ")
		input, _ := r.ReadString('\n')
		input = strings.TrimSpace(input)
		number, err := strconv.Atoi(input)
		if err != nil {
			cleanScreen()
			fmt.Println("Invalid choice")
			continue
		}

		if item, ok := menu.Menu[menu.MenuOption(number)]; ok {
			cleanScreen()
			fmt.Printf("You selected: %s\n", item.Name)
			item.Action()

			// Wait for user to press enter
			fmt.Print("Press 'Enter' to continue...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')

		} else {
			cleanScreen()
			fmt.Println("Invalid choice")
			menu.PrintMenu()
		}
	}
}
