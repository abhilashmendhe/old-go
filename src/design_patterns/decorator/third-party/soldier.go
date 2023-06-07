package thirdparty

import "log"

type InterfaceSoldier interface {
	Attack() int
	Defense() int
}

type BasicSoldier struct {
}

func (b BasicSoldier) Attack() int {
	return 1
}
func (b BasicSoldier) Defense() int {
	return 1
}

func DisplayStats(soldier InterfaceSoldier) {
	log.Printf("soldier stats: attack %d, defense %d", soldier.Attack(), soldier.Defense())
}
