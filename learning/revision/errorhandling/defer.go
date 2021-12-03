package errorhandling

import "fmt"

func DeferMultiple() {

	multipleDefers()
	//Note:
	// - Multiple defers are called in line-decrement order
}

func multipleDefers() {
	defer thisWillBeCalledLast()
	defer thisWillBeCalledLastButOne()
	fmt.Println("First Call this, then go to defer")
}

func thisWillBeCalledLastButOne() {
	fmt.Println("defer: thisWillBeCalledLastButOne")
}

func thisWillBeCalledLast() {
	fmt.Println("defer: thisWillBeCalledLast")
}

