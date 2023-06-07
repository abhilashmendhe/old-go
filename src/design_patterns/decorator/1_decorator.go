/*
Wrapper or decorator patterns
Decorator is a structural design pattern that lets you attach new behaviors to objects
by placing these objects inside special wrapper objects that contain the behaviors.

Pros: 1)
*/

package main

import thirdparty "design_patterns/decorator/third-party"

func main() {
	basicSoldier := thirdparty.BasicSoldier{}
	thirdparty.DisplayStats(basicSoldier)

	soldierWithSword := SoldierWithSword{soldier: basicSoldier}
	thirdparty.DisplayStats(soldierWithSword)

	soldierWithShield := SoldierWithSheild{soldier: basicSoldier}
	thirdparty.DisplayStats(soldierWithShield)

	soldierWithShieldWithSword := SoldierWithSheild{
		soldier: SoldierWithSword{
			soldier: basicSoldier,
		},
	}
	thirdparty.DisplayStats(soldierWithShieldWithSword)
}

// Decorator 1: Soldier with a sword
type SoldierWithSword struct {
	soldier thirdparty.InterfaceSoldier
}

func (s SoldierWithSword) Attack() int {
	return s.soldier.Attack() + 10
}
func (s SoldierWithSword) Defense() int {
	return s.soldier.Defense()
}

// Decorator 2: Soldier with a shield
type SoldierWithSheild struct {
	soldier thirdparty.InterfaceSoldier
}

func (s SoldierWithSheild) Attack() int {
	return s.soldier.Attack() + 5
}
func (s SoldierWithSheild) Defense() int {
	return s.soldier.Defense() + 10
}
