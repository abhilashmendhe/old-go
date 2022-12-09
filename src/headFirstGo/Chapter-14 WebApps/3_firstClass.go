package main

import "fmt"

func hi() {
	fmt.Println("Hi")
}

func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

func doMath(passedFunction func(int, int) float64) {
	res := passedFunction(10, 20)
	fmt.Println(res)
}

func div(a int, b int) float64 {
	return float64(a) / float64(b)
}

func mult(a int, b int) float64 {
	return float64(a) * float64(b)
}

func main() {
	var greeterFunction func()
	var mathFunction func(int, int) float64
	mathFunction = divide
	greeterFunction = hi
	greeterFunction()
	fmt.Println(mathFunction(10, 21))

	doMath(div)
	doMath(mult)
}
