package snippets

import "fmt"

func Stacks() {
	push(1)
	push(2)
	push(3)
	push(4)
	fmt.Println("-------------------")
	fmt.Println("stack : ", globalStack.stackArray)
	fmt.Println("top : ", globalStack.top)
	fmt.Println("-------------------")
	pop()
	fmt.Println("-------------------")
	fmt.Println("stack : ", globalStack.stackArray)
	fmt.Println("top : ", globalStack.top)
	fmt.Println("-------------------")
	pop()
	fmt.Println("-------------------")
	fmt.Println("stack : ", globalStack.stackArray)
	fmt.Println("top : ", globalStack.top)
	fmt.Println("-------------------")
	pop()
	fmt.Println("-------------------")
	fmt.Println("stack : ", globalStack.stackArray)
	fmt.Println("top : ", globalStack.top)
	fmt.Println("-------------------")
	pop()
	fmt.Println("-------------------")
	fmt.Println("stack : ", globalStack.stackArray)
	fmt.Println("top : ", globalStack.top)
	fmt.Println("-------------------")
	pop()
	fmt.Println("-------------------")
	fmt.Println("stack : ", globalStack.stackArray)
	fmt.Println("top : ", globalStack.top)
	fmt.Println("-------------------")
}

var globalStack stack

type stack struct {
	top        int
	stackArray []int
}

func push(value int) {
	tempValue := append(globalStack.stackArray, value)
	globalStack = stack{top: value, stackArray: tempValue}
}

func pop() {
	length := len(globalStack.stackArray)
	if length != 0 {
		if length == 1 {
			tempArr := []int{}
			globalStack = stack{top: 0, stackArray: tempArr}
		} else {
			tempArr := globalStack.stackArray[:length-1]
			topValue := tempArr[len(tempArr)-1]
			globalStack = stack{top: topValue, stackArray: tempArr}
		}

	} else {
		fmt.Println("Empty stack!!")
	}
}
