package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func cleanScreen() {
	fmt.Print("\033[H\033[2J")
}

func readIntFromConsole(reader *bufio.Reader) (int, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	// Trim the newline character from the input
	input = strings.TrimSpace(input)

	// Parse the string into an integer
	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	fmt.Println()

	return number, nil
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	cleanScreen()

	printMenu()
	initLaunchers()

	for {
		fmt.Print("Enter your choice: ")
		choice, _ := readIntFromConsole(reader)

		if item, ok := menu[menuOption(choice)]; ok {
			cleanScreen()
			fmt.Printf("You selected: %s\n", item.name)
			item.action()
		} else {
			cleanScreen()
			fmt.Println("Invalid choice")
			printMenu()
		}
	}
}
