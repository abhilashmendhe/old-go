package main

import "fmt"

// for constant
const scalar = 0.1 // if you want to export then you just capatilize the first letter of the variable name

// variable grouping
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

// functions that panic
func parseIntFromString(s string) (int, error) {
	// Logic

	return 10, nil
}

func main() {

}
