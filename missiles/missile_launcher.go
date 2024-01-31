package missiles

import (
	"fmt"
	"math/rand"
	"sort"
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func removeMissilesByIndexes(missiles []missile, indexes []int) []missile {
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

type MissileLauncher interface {
	Launch(count int) int
	Add(count int)
	Clear()
	ClearAt(index int)
	Len() int
}

type missile struct {
	failed bool
}

type missileStorage struct {
	missiles []missile
}

func (m *missileStorage) Add(count int) {
	for i := 0; i < count; i++ {
		m.missiles = append(m.missiles, missile{
			failed: false,
		})
	}
}

func (m *missileStorage) ClearAt(index int) {
	if index < len(m.missiles) && index >= 0 { // Check for valid index
		m.missiles = append(m.missiles[:index], m.missiles[index+1:]...)
	}
}

func (m *missileStorage) Clear() {
	m.missiles = []missile{}
}

func (m *missileStorage) Len() int {
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

func (t *torpedoMissileLauncher) Launch(count int) int {
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

	t.missiles = removeMissilesByIndexes(t.missiles, launchedIndexes)
	return launched
}

type ballisticMissileLauncher struct {
	*missileStorage
	successRate int
}

func (b *ballisticMissileLauncher) Launch(count int) int {
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

	b.missiles = removeMissilesByIndexes(b.missiles, launchedIndexes)

	return launched
}

type cruiseMissileLauncher struct {
	*missileStorage
	successRate int
}

func (c *cruiseMissileLauncher) Launch(count int) int {
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

	c.missiles = removeMissilesByIndexes(c.missiles, launchedIndexes)
	return launched
}

type hypersonicMissileLauncher struct {
	*missileStorage
	maxRange int
}

func (h *hypersonicMissileLauncher) Launch(count int) int {
	launched := 0
	launchedIndexes := make([]int, 0)

	distance := random(0, h.maxRange)

	for {
		fmt.Print("How far would you like to launch: ")

		if distance > h.maxRange {
			fmt.Println("This is too far, the maximum range is", h.maxRange)
		} else {
			break
		}
	}

	successRate := (1.0 - (float64(distance) / float64(h.maxRange))) * 100.0
	for i := 0; i < count; i++ {
		m := h.missiles[i]

		if m.failed {
			continue
		}

		missileHitRate := random(0, 100)
		if missileHitRate < int(successRate) {
			launchedIndexes = append(launchedIndexes, i)
			launched++
		} else {
			m.failed = false
		}
	}

	h.missiles = removeMissilesByIndexes(h.missiles, launchedIndexes)
	return launched
}

var Launchers = map[Launcher]MissileLauncher{
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
	hypersonicLauncher: &hypersonicMissileLauncher{
		missileStorage: newMissileStorage(),
		maxRange:       1500,
	},
}

func InitLaunchers() {
	for _, l := range Launchers {
		l.Add(10)
	}
}
