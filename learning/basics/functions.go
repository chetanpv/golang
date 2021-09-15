package basics

import (
	"fmt"
)

// Functions shows how to call a function in differnt ways
func Functions() {
	// Calling a function
	simpleFunction()

	// Calling a function with return value
	name := funtionWithReturnValue()
	fmt.Println(name)

	// Calling a function with return values
	middleName, lastName := functionWithMultipleReturnValues()
	fmt.Println(middleName, lastName)

	// Passing arguement to functions
	value := functionWithArgs(3)
	fmt.Println(value)

	//------------------------------ Empty interface OR passing arg without type ------------------------------
	functionWIthInterfaces("Chetan")
	functionWIthInterfaces(0)
	functionWIthInterfaces(true)

	// Variadic functions
	funtionWithDynamicArgs("a", "b", "c")
}

func simpleFunction() {
	fmt.Println("Hello")
}

func funtionWithReturnValue() string {
	return "Chetan"
}

func functionWithMultipleReturnValues() (string, string) {
	return "Chetan", "Patige"
}

func functionWithArgs(number int) int {
	return 1 + number
}

func funtionWithDynamicArgs(characters ...string) {
	fmt.Println(characters) // ["a","b","c"]
}

func functionWIthInterfaces(a interface{}) {
	fmt.Println(a)
}
