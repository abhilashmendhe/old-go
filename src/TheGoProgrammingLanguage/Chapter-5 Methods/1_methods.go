package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int32
}

type Path []Point

func (p Point) Distance(q Point) float64 {
	x := (p.X - q.X) * (p.X - q.X)
	y := (p.Y - q.Y) * (p.Y - q.Y)
	return math.Sqrt(float64(x) + float64(y))
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	perim := Path{
		{1, 2},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(perim.Distance())
}
