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

var menu map[menuOption]menuItem = map[menuOption]menuItem{
	menuStoreNewMissles: {
		name: "Store new missiles",
		action: func() {
			println("Storing new missiles")
		},
	},
	menuLaunchMissle: {
		name: "Launch missile",
		action: func() {
			println("Launching missile")
		},
	},
	menuInventoryReport: {
		name: "Inventory report",
		action: func() {
			println("Inventory report")
		},
	},
	menuCleanOutMissles: {
		name: "Clean out missiles",
		action: func() {
			println("Cleaning out missiles")
		},
	},
	menuShutdown: {
		name: "Shutdown",
		action: func() {
			fmt.Println("GG WP, exiting.....")
			os.Exit(0)
		},
	},
}

func printMenu() {
	for i := 1; i <= len(menu); i++ {
		fmt.Printf("%d. %s\n", i, menu[menuOption(i)].name)
	}
}
