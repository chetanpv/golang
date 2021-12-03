package revision

import "fmt"

func CheckType() {
	var a interface{}
	a = "Chetan"
	switch b:= a.(type) {
	case int:
		fmt.Println("This is int", b)
	case string:
		fmt.Println("This is string", b)
	}
}
