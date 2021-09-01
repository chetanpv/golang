package basics

import "fmt"

// Conditions ...
func Conditions() {
	ifBlock()
	switchBlock()
}

func ifBlock() {
	if 5 > 3 {
		fmt.Println("True")
	} else if 3 > 5 {
		fmt.Println("True")
	} else {
		fmt.Println("True")
	}
}

func switchBlock() {
	i := 2
	switch i {
	case 1:
		fmt.Println("Its 1")
	case 2:
		callFunc()
	default:
		fmt.Println("No matches")
	}
}

func callFunc() {
	fmt.Println("Got from switch")
}
