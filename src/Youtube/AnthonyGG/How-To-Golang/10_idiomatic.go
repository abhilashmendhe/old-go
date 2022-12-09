package main

import (
	"fmt"
	"net"
	"sync"
)

// 1) for constant
const scalar = 0.1 // if you want to export then you just capatilize the first letter of the variable name

// 2) variable grouping
func Foo() int {

	// variable grouping
	// x := 100
	// y := 2
	// foo := "foo"
	// instead do
	var (
		x   = 100
		y   = 2
		foo = "foo"
	)

	fmt.Println(x, y, foo)
	return x + y
}

// 3) functions that panic, if there is a panic then always prefix with "Must" to your function name
func MustParseIntFromString(s string) (int, error) {
	// Logic
	panic("oops")
	return 10, nil
}

// 4) how to struct initialize
type Vector struct {
	x int
	y int
}

// 5) Mutex grouping
type Server struct {
	listenAddress string
	isRunning     bool

	// here the "map" type we want to protect therefore we always put mutex above the type that we want to protect...
	// for mutex the variable name should alwyas contain a suffix ("mu" or "Lock")
	mu    sync.RWMutex
	peers map[string]net.Conn
}

// 6) interface declarations/naming, you should add a suffix "er"
type Storager interface {
	Get()
	Put()
	Delete()
	Patch()
}

type Getter interface {
	Get()
}

type Putter interface {
	Put()
}

// 7) function grouping, so always put your exported functino at the top of your file and then all private function
// to be placed below

func VeryImportantFuncExported() {} // exported
func veryImportantFunc()         {} // unexported

// 8) http handler naming, "handle" with prefix
func handleGetUserById() {}
func handleResizeImage() {}

// 9) enum

type Suit byte

// you should add the prefix type name to all the const variable, like below you can see Suit(...), Suit is type..
const (
	SuitHearts Suit = iota
	SuitClubs
	SuitDiamonds
	SuitSpades
)

// 10 constructor

// we want to construct this Order
type Order struct {
	Size float64
}

func (o Order) NewOrder(size float64) *Order {
	// logic here
	return &Order{Size: size}
}
func main() {

	//4) this is how you initialize with name variables....
	vec := Vector{
		x: 10,
		y: 20,
	}
	fmt.Println(vec)
	oo := &Order{Size: 10.2}
	oo.NewOrder(100)
}
