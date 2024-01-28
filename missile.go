package main

import (
	"fmt"
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
	launch()
	addMissiles(count int)
	clearMissiles()
	len() int
}

type missile struct {
	failed bool
}

type missileStorage struct {
	missiles []missile
}

func (m *missileStorage) addMissiles(count int) {
	for i := 0; i < count; i++ {
		m.missiles = append(m.missiles, missile{
			failed: false,
		})
	}
}

func (m *missileStorage) clearMissiles() {
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
}

func (t *torpedoMissileLauncher) launch() {
	println("Torpedo launched!")
}

type ballisticMissileLauncher struct {
	*missileStorage
}

func (b *ballisticMissileLauncher) launch() {
	println("Ballistic missile launched!")
}

type cruiseMissileLauncher struct {
	*missileStorage
}

func (c *cruiseMissileLauncher) launch() {
	println("Cruise missile launched!")
}

type hypersonicMissileLauncher struct {
	*missileStorage
}

func (h *hypersonicMissileLauncher) launch() {
	println("Hypersonic missile launched!")
}

var launchers = map[launcher]missileLauncher{
	torpedoLauncher: &torpedoMissileLauncher{
		missileStorage: newMissileStorage(),
	},
	ballisticLauncher: &ballisticMissileLauncher{
		missileStorage: newMissileStorage(),
	},
	cruiseLauncher: &cruiseMissileLauncher{
		missileStorage: newMissileStorage(),
	},
	hypersonicLauncher: &hypersonicMissileLauncher{
		missileStorage: newMissileStorage(),
	},
}

func initLaunchers() {
	for _, l := range launchers {
		l.addMissiles(10)
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
}
