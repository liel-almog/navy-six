package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/lielalmog/navy-six/missiles"
)

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

		missileLauncher.Add(missilesCount)
		fmt.Printf("Added %d missiles to %s launcher\n", missilesCount, launcherType)
		break
	}
}

func launchAllMissiles() {
	sl := make(map[missiles.Launcher]int)
	var total int

	for lt, ml := range missiles.Launchers {
		s := ml.Launch(ml.Len())
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
	var launcherType missiles.Launcher
	var mLauncher missiles.MissileLauncher

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

		if missiles.IsLauncher(number) {
			launcherType = missiles.Launcher(number)
			mLauncher = missiles.Launchers[launcherType]
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

		if missilesCount > mLauncher.Len() {
			mLauncher.Add(missilesCount - mLauncher.Len())
		}

		successfulLaunches := mLauncher.Launch(missilesCount)
		fmt.Printf("Launched %d missiles successfully from %s launcher\n", successfulLaunches, launcherType)
		break
	}
}

func inventoryReport() {
	launcherTypes := missiles.OrderedLaunchers()
	var totalMissiles int

	t := table.NewWriter()

	t.SetAutoIndex(true)
	t.AppendHeader(table.Row{"Launcher Type", "Missiles"})

	for _, lt := range launcherTypes {
		ml := missiles.Launchers[lt]
		totalMissiles += ml.Len()
		t.AppendRow(table.Row{lt.String(), ml.Len()})
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
			for _, ml := range missiles.Launchers {
				ml.Clear()
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
		for _, ml := range missiles.Launchers {
			if indexToRemove < ml.Len() {
				ml.ClearAt(indexToRemove)
				fmt.Printf("Missile at index %d cleaned out\n", indexToRemove)
				return
			}

			totalMissiles += ml.Len()
			indexToRemove -= ml.Len()
		}

		fmt.Printf("Invalid index, we have a total of %d missiles. Please try again\n", totalMissiles)
	}
}

func shutdown() {
	fmt.Println("GG WP, exiting.....")
	os.Exit(0)
}
