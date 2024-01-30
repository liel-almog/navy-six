package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
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
	sl := make(map[launcher]int)
	var total int

	for lt, ml := range launchers {
		s := ml.launch(ml.len())
		sl[lt] = s
		total += s
	}

	t := table.NewWriter()

	t.SetAutoIndex(true)
	t.AppendHeader(table.Row{"Launcher Type", "Successful Launches"})
	t.AppendFooter(table.Row{"Total", total})

	for lt, s := range sl {
		t.AppendRow(table.Row{lt.String(), s})
	}

	t.SetCaption("Missile Total War Launch Report")
	fmt.Println(t.Render())
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
			mLauncher.add(missilesCount - mLauncher.len())
		}

		successfulLaunches := mLauncher.launch(missilesCount)
		fmt.Printf("Launched %d missiles successfully from %s launcher\n", successfulLaunches, launcherType)
		break
	}
}

func inventoryReport() {
	launcherTypes := orderLaunchers()
	var totalMissiles int

	t := table.NewWriter()

	t.SetAutoIndex(true)
	t.AppendHeader(table.Row{"Launcher Type", "Missiles"})

	for _, lt := range launcherTypes {
		ml := launchers[lt]
		totalMissiles += ml.len()
		t.AppendRow(table.Row{lt.String(), ml.len()})
	}

	t.AppendFooter(table.Row{"Total", totalMissiles})
	t.SetCaption("Missile Inventory Report")

	fmt.Println(t.Render())
}

func clearMissiles() {
	const cleanAll = "All"
	var totalMissiles int
	r := bufio.NewReader(os.Stdin)

	// Clean out missiles only if the input was not a number
	for {

		// cleanScreen()
		// Write to the terminal instread of the user
		fmt.Printf("To clean out all missiles, type '%s'\n", cleanAll)
		fmt.Println("To clean out a missile at a specific index, type the index number")
		fmt.Print("Enter your choice: ")
		// Write to the teminal so the reader can read it

		input, _ := r.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == cleanAll {
			for _, ml := range launchers {
				ml.clear()
			}

			fmt.Println("All missiles cleaned out")
			break
		}

		indexToRemove, err := convertStringToInt(input)
		if err != nil {
			fmt.Println("Invalid input, please try again")
			continue
		}

		if indexToRemove < 0 {
			fmt.Println("Please enter a positive number")
			continue
		}

		if indexToRemove == 0 {
			break
		}

		// Clean out missiles at a specific index
		for _, ml := range launchers {
			if indexToRemove < ml.len() {
				ml.clearAt(indexToRemove)
				fmt.Printf("Missile at index %d cleaned out\n", indexToRemove)
				return
			}

			totalMissiles += ml.len()
			indexToRemove -= ml.len()
		}

		fmt.Printf("Invalid index, we have a total of %d missiles. Please try again\n", totalMissiles)
	}
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
		action: clearMissiles,
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
