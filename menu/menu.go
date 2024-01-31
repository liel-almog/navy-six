package menu

type MenuOption int

const (
	menuStoreNewMissles MenuOption = iota + 1
	menuLaunchMissle
	menuInventoryReport
	menuCleanOutMissles
	menuShutdown
)

type MenuItem struct {
	Name   string
	Action func()
}

var Menu map[MenuOption]MenuItem = map[MenuOption]MenuItem{
	menuStoreNewMissles: {
		Name:   "Store new missiles",
		Action: storeNewMissiles,
	},
	menuLaunchMissle: {
		Name:   "Launch missile",
		Action: launchMissile,
	},
	menuInventoryReport: {
		Name:   "Inventory report",
		Action: inventoryReport,
	},
	menuCleanOutMissles: {
		Name:   "Clean out missiles",
		Action: clearMissiles,
	},
	menuShutdown: {
		Name:   "Shutdown",
		Action: shutdown,
	},
}
