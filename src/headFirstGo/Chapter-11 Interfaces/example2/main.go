package main

import (
	"fmt"
	"headFirstGo/interfaces/example2/mypkg"
)

func main() {
	var value mypkg.MyInterface

	value = mypkg.MyType(5)
	value.MethodWithParameters(10.21)
	value.MethodWithoutParameters()
	s := value.MethodWithReturnValue()
	fmt.Println(s)

	var val mypkg.MyType
	val.MethodNotInInterface()
	val.MethodWithParameters(2.1)
	// value.MethodNotInInterface()

	var tval mypkg.MyT2
	tval.MethodWithoutParameters()
}
