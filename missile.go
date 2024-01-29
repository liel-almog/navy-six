package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

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
	successCount := 0

	for i := 0; i < count; i++ {
		m := t.missiles[i]

		if m.failed {
			continue
		}

		missileHitRate := random(0, 100)
		if missileHitRate < t.successRate {
			successCount++
		} else {
			m.failed = true
		}
	}

	return successCount
}

type ballisticMissileLauncher struct {
	*missileStorage
	successRate int
}

func (b *ballisticMissileLauncher) launch(count int) int {
	successCount := 0

	for i := 0; i < count; i++ {
		m := b.missiles[i]
		if m.failed {
			continue
		}

		missileHitRate := random(0, 100)
		if missileHitRate < b.successRate {
			successCount++
		} else {
			m.failed = true
		}
	}

	return successCount
}

type cruiseMissileLauncher struct {
	*missileStorage
	successRate int
}

func (c *cruiseMissileLauncher) launch(count int) int {
	successCount := 0

	for i := 0; i < count; i++ {
		m := c.missiles[i]

		if m.failed {
			continue
		}

		missileHitRate := random(0, 100)
		if missileHitRate < c.successRate {
			successCount++
		} else {
			m.failed = true
		}
	}

	return successCount
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

func printMissilesLaunchers() {
	var launcherTypes []launcher

	for key := range launchers {
		launcherTypes = append(launcherTypes, key)
	}

	// sort the launchers
	slices.Sort(launcherTypes)

	for i := 0; i < len(launcherTypes); i++ {
		fmt.Printf("%d. %s\n", i+1, launcherTypes[i])
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
