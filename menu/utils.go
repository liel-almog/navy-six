package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lielalmog/navy-six/missiles"
)

func printMissilesLaunchers() {
	launchers := missiles.OrderedLaunchers()

	for i := 0; i < len(launchers); i++ {
		fmt.Printf("%d. %s\n", i+1, launchers[i])
	}

	fmt.Println()
}

func PrintMenu() {
	fmt.Println("Menu:")
	for i := 1; i <= len(Menu); i++ {
		fmt.Printf("%d. %s\n", i, Menu[MenuOption(i)].Name)
	}

	fmt.Println()
}

func selectMissleLauncher() (missiles.Launcher, missiles.MissileLauncher) {
	r := bufio.NewReader(os.Stdin)

	for {
		cleanScreen()
		fmt.Println("Please select a launcher:")
		printMissilesLaunchers()
		fmt.Print("Selected launcher: ")

		number, _ := readIntFromConsole(r)

		if !missiles.IsLauncher(number) {
			fmt.Println("Invalid input, please try again")
			continue
		}

		launcherType := missiles.Launcher(number)
		return launcherType, missiles.Launchers[launcherType]
	}
}

func convertStringToInt(s string) (int, error) {
	s = strings.TrimSpace(s)

	// Parse the string into an integer
	number, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	fmt.Println()

	return number, nil
}

func readIntFromConsole(r *bufio.Reader) (int, error) {
	input, err := r.ReadString('\n')
	if err != nil {
		return 0, err
	}

	number, err := convertStringToInt(input)
	if err != nil {
		return 0, err
	}

	return number, nil
}

func cleanScreen() {
	fmt.Print("\033[H\033[2J")
}
