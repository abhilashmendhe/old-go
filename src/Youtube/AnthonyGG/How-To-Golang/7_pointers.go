package main

import (
	"fmt"
)

type BigData struct {
	// 500 MB

}

func processBigData(bd *BigData) {

}

type Player struct {
	health int
}

func (player *Player) takeDamageFromExplosion(dmg int) {
	fmt.Println("player is taking damage from explosion")
	player.health -= dmg
}
func takeDamageFromExplosion(player Player, explosion_dmg int) { //passing 8byte pointer struct
	fmt.Println("player is taking damage from explosion")
	player.health -= explosion_dmg
}

func main() {

	player := Player{ // this works without puting & before Player (&Player).. only works when calling method not functions.
		health: 100,
	} // 8 byte long int
	// player = nil
	fmt.Printf("Before explosion %+v\n", player)
	player.takeDamageFromExplosion(80)
	fmt.Printf("After explosion %+v\n", player)

}
