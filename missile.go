package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func removeByIndexes(missiles []missile, indexes []int) []missile {
	// Sort the indexes in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(indexes)))

	// Remove the elements by index
	for _, index := range indexes {
		if index < len(missiles) && index >= 0 { // Check for valid index
			missiles = append(missiles[:index], missiles[index+1:]...)
		}
	}

	return missiles
}

type launcher int

const (
	torpedoLauncher launcher = iota + 1
	ballisticLauncher
	cruiseLauncher
	hypersonicLauncher
)

func (l launcher) String() string {
	switch l {
	case torpedoLauncher:
		return "Torpedo"
	case ballisticLauncher:
		return "Ballistic"
	case cruiseLauncher:
		return "Cruise"
	case hypersonicLauncher:
		return "Hypersonic"
	default:
		return "Unknown"
	}
}

type missileLauncher interface {
	// Launches the missiles and returns the number of missiles that successfully launched
	launch(count int) int
	add(count int)
	clear()
	len() int
}

type missile struct {
	failed bool
}

type missileStorage struct {
	missiles []missile
}

func (m *missileStorage) add(count int) {
	for i := 0; i < count; i++ {
		m.missiles = append(m.missiles, missile{
			failed: false,
		})
	}
}

func (m *missileStorage) clear() {
	m.missiles = []missile{}
}

func (m *missileStorage) len() int {
	return len(m.missiles)
}

func newMissileStorage() *missileStorage {
	return &missileStorage{
		missiles: []missile{},
	}
}

type torpedoMissileLauncher struct {
	*missileStorage
	successRate int
}

func (t *torpedoMissileLauncher) launch(count int) int {
	launched := 0
	launchedIndexes := make([]int, 0)

	for i := 0; i < count; i++ {
		m := t.missiles[i]

		if m.failed {
			continue
		}

		missileHitRate := random(0, 100)
		if missileHitRate < t.successRate {
			launched++
			launchedIndexes = append(launchedIndexes, i)
		} else {
			m.failed = true
		}
	}

	t.missiles = removeByIndexes(t.missiles, launchedIndexes)
	return launched
}

type ballisticMissileLauncher struct {
	*missileStorage
	successRate int
}

func (b *ballisticMissileLauncher) launch(count int) int {
	launched := 0
	launchedIndexes := make([]int, 0)

	for i := 0; i < count; i++ {
		m := b.missiles[i]
		if m.failed {
			continue
		}

		missileHitRate := random(0, 100)
		if missileHitRate < b.successRate {
			launched++
			launchedIndexes = append(launchedIndexes, i)
		} else {
			m.failed = true
		}
	}

	b.missiles = removeByIndexes(b.missiles, launchedIndexes)

	return launched
}

type cruiseMissileLauncher struct {
	*missileStorage
	successRate int
}

func (c *cruiseMissileLauncher) launch(count int) int {
	launched := 0
	launchedIndexes := make([]int, 0)

	for i := 0; i < count; i++ {
		m := c.missiles[i]

		if m.failed {
			continue
		}

		missileHitRate := random(0, 100)
		if missileHitRate < c.successRate {
			launched++
			launchedIndexes = append(launchedIndexes, i)
		} else {
			m.failed = true
		}
	}

	c.missiles = removeByIndexes(c.missiles, launchedIndexes)
	return launched
}

var launchers = map[launcher]missileLauncher{
	torpedoLauncher: &torpedoMissileLauncher{
		missileStorage: newMissileStorage(),
		successRate:    100,
	},
	ballisticLauncher: &ballisticMissileLauncher{
		missileStorage: newMissileStorage(),
		successRate:    50,
	},
	cruiseLauncher: &cruiseMissileLauncher{
		missileStorage: newMissileStorage(),
		successRate:    20,
	},
}

func initLaunchers() {
	for _, l := range launchers {
		l.add(10)
	}
}

func orderLaunchers() []launcher {
	var launcherTypes []launcher

	for key := range launchers {
		launcherTypes = append(launcherTypes, key)
	}

	// sort the launchers
	slices.Sort(launcherTypes)

	return launcherTypes
}

func printMissilesLaunchers() {
	launchers := orderLaunchers()

	for i := 0; i < len(launchers); i++ {
		fmt.Printf("%d. %s\n", i+1, launchers[i])
	}

	fmt.Println()
}

func selectMissleLauncher() (launcher, missileLauncher) {
	r := bufio.NewReader(os.Stdin)

	for {
		cleanScreen()
		fmt.Println("Please select a launcher:")
		printMissilesLaunchers()
		fmt.Print("Selected launcher: ")

		number, _ := readIntFromConsole(r)

		if !isLauncher(number) {
			fmt.Println("Invalid input, please try again")
			continue
		}

		launcherType := launcher(number)
		return launcherType, launchers[launcherType]
	}
}

func isLauncher(l int) bool {
	_, ok := launchers[launcher(l)]
	return ok
}
