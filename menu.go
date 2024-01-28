package main

import (
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
	for i := 1; i <= len(menu); i++ {
		fmt.Printf("%d. %s\n", i, menu[menuOption(i)].name)
	}
}
