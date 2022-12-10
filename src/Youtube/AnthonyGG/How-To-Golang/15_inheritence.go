package main

import "fmt"

type SpecialPosition struct {
	Position
}

func (sp *SpecialPosition) MoveSpecial(x, y float64) {
	sp.x += x * x
	sp.y += y * y
}

type Position struct {
	x float64
	y float64
}

func (p *Position) Move(x, y float64) {
	p.x += x
	p.y += y
}

func (p *Position) Teleport(x, y float64) {
	p.x = x
	p.y = y
}

func NewPlayerr() *Playerr {
	return &Playerr{
		Position: &Position{},
	}
}

type Playerr struct {
	*Position
}

type Enemy struct {
	*SpecialPosition
}

func NewEnemey() *Enemy {
	return &Enemy{
		SpecialPosition: &SpecialPosition{},
	}
}

func main() {

	player := NewPlayerr()
	player.Move(100, 21)
	fmt.Println(player.Position)
	player.Teleport(2012, 1231)
	fmt.Println(player.Position)

	boss := NewEnemey()
	boss.Move(1.1, 21)
	fmt.Println(boss.Position)
	boss.MoveSpecial(1.1, 2.1)
	fmt.Println(boss.SpecialPosition)

}
