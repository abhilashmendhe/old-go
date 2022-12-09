package main

import "encaps/example2/mypackage"

func main() {
	value := mypackage.MyType{}
	value.ExportedMethod()
	// value.unexportedMethod() // won't be exported
}
