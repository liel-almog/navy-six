package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	r := bufio.NewReader(os.Stdin)

	launcherType, missileLauncher := selectMissleLauncher()
	for {
		fmt.Println("How many missiles you want to add to", launcherType, "Launcher")
		fmt.Print("Enter number of missiles: ")
		missilesCount, err := readIntFromConsole(r)

		if err != nil {
			fmt.Println("Invalid input, please try again")
			continue
		}

		if missilesCount < 0 {
			fmt.Println("Please enter a positive number")
			continue
		}

		missileLauncher.add(missilesCount)
		fmt.Printf("Added %d missiles to %s launcher\n", missilesCount, launcherType)
		break
	}
}

func launchAllMissiles() {
	var sl map[launcher]int = make(map[launcher]int)

	for lt, ml := range launchers {
		s := ml.launch(ml.len())
		sl[lt] = s
	}

	var total int
	for lt, s := range sl {
		fmt.Printf("Launched %d missiles from %s launcher\n", s, lt)
		total += s
	}

	fmt.Println("Launched all missiles")
}

func launchMissile() {
	const totalWar = "TotalWar"
	var launcherType launcher
	var mLauncher missileLauncher

	r := bufio.NewReader(os.Stdin)

	for {
		cleanScreen()
		fmt.Println("Please select a launcher:")
		printMissilesLaunchers()
		fmt.Print("Selected launcher: ")

		input, _ := r.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == totalWar {
			launchAllMissiles()
			return
		}

		number, err := convertStringToInt(input)
		if err != nil {
			fmt.Println("Invalid input, please try again")
			continue
		}

		if isLauncher(number) {
			launcherType = launcher(number)
			mLauncher = launchers[launcherType]
			break
		}
	}

	for {
		fmt.Println("How many missiles you want to launch from", launcherType, "Launcher")
		fmt.Print("Enter number of missiles: ")
		missilesCount, err := readIntFromConsole(r)

		if err != nil {
			fmt.Println("Invalid input, please try again")
			continue
		}

		if missilesCount < 0 {
			fmt.Println("Please enter a positive number")
			continue
		}

		if missilesCount > mLauncher.len() {
			mLauncher.add(missilesCount)
		}

		successfulLaunches := mLauncher.launch(missilesCount)
		fmt.Printf("Launched %d missiles successfully from %s launcher\n", successfulLaunches, launcherType)
		break
	}
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
