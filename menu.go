package main

import (
	"fmt"
	"os"
)

type menuItem struct {
	name   string
	action func()
}

var menu map[int]menuItem

func storeNewMissles() {
	fmt.Println("Add")
}

func launchMissle() {
	fmt.Println("Subtract")
}

func inventoryReport() {
	fmt.Println("Multiply")
}

func cleanOutMissles() {
	fmt.Println("Divide")
}

func shutdown() {
	fmt.Println("GG WP, exiting.....")
	os.Exit(0)
}

func initMenu() {
	menu = make(map[int]menuItem)
	menu[1] = menuItem{name: "Store new missles", action: storeNewMissles}
	menu[2] = menuItem{name: "Launch missle", action: launchMissle}
	menu[3] = menuItem{name: "Inventory report", action: inventoryReport}
	menu[4] = menuItem{name: "Clean out missles", action: cleanOutMissles}
	menu[5] = menuItem{name: "Shutdown", action: shutdown}
}
