package mypkg

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameters(float64)
	MethodWithReturnValue() string
}

type MyType int
type MyT2 float64

func (m MyT2) MethodWithoutParameters() {
	fmt.Println("Method without params...")
}

func (m MyType) MethodWithoutParameters() {
	fmt.Println("Method without params...")
}

func (m MyType) MethodWithParameters(f float64) {
	fmt.Println("Method with params:", f)
}

func (m MyType) MethodWithReturnValue() string {
	return "From return value method"
}

func (m MyType) MethodNotInInterface() {
	fmt.Println("Method outside interface.!.!.")
}
