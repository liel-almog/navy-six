package missiles

import "slices"

type Launcher int

const (
	torpedoLauncher Launcher = iota + 1
	ballisticLauncher
	cruiseLauncher
	hypersonicLauncher
)

func (l Launcher) String() string {
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

func OrderedLaunchers() []Launcher {
	var launcherTypes []Launcher

	for key := range Launchers {
		launcherTypes = append(launcherTypes, key)
	}

	slices.Sort(launcherTypes)

	return launcherTypes
}

func IsLauncher(l int) bool {
	_, ok := Launchers[Launcher(l)]
	return ok
}
