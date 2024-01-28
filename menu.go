package main

import (
	"bufio"
	"fmt"
	"os"
)

type menuOption int

const (
	menuStoreNewMissles menuOption = iota + 1
	menuLaunchMissle
	menuInventoryReport
	menuCleanOutMissles
	menuShutdown
)

type menuItem struct {
	name   string
	action func()
}

func storeNewMissiles() {
	reader := bufio.NewReader(os.Stdin)

	printMissilesLaunchers()
	for {
		fmt.Print("To which launcher you want to add missiles: ")
		number, _ := readIntFromConsole(reader)
		launcherType := launcher(number)

		if _, ok := launchers[launcherType]; ok {
			for {
				fmt.Println("How many missiles you want to add to", launcherType, "Launcher")
				fmt.Print("Enter number of missiles: ")
				missilesCount, err := readIntFromConsole(reader)

				if err != nil {
					fmt.Println("Invalid input, please try again")
					continue
				}

				if missilesCount < 0 {
					fmt.Println("Please enter a positive number")
					continue
				}

				launchers[launcherType].addMissiles(missilesCount)
				fmt.Printf("Added %d missiles to %s launcher\n", missilesCount, launcherType)
				break
			}

			break
		} else {
			cleanScreen()
			fmt.Println("Please select again")
			printMissilesLaunchers()
		}
	}

}

func launchMissile() {
	fmt.Println("Launch missile")
}

func inventoryReport() {
	fmt.Println("Inventory report")
}

func cleanOutMissiles() {
	fmt.Println("Clean out missiles")
}

func shutdown() {
	fmt.Println("GG WP, exiting.....")
	os.Exit(0)
}

var menu map[menuOption]menuItem = map[menuOption]menuItem{
	menuStoreNewMissles: {
		name:   "Store new missiles",
		action: storeNewMissiles,
	},
	menuLaunchMissle: {
		name:   "Launch missile",
		action: launchMissile,
	},
	menuInventoryReport: {
		name:   "Inventory report",
		action: inventoryReport,
	},
	menuCleanOutMissles: {
		name:   "Clean out missiles",
		action: cleanOutMissiles,
	},
	menuShutdown: {
		name:   "Shutdown",
		action: shutdown,
	},
}

func printMenu() {
	fmt.Println("Menu:")
	for i := 1; i <= len(menu); i++ {
		fmt.Printf("%d. %s\n", i, menu[menuOption(i)].name)
	}

	fmt.Println()
}
