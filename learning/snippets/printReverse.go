package snippets

import "fmt"

func ReversePrint() {
	a := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(a)
	for i := len(a) - 1; i <= 0; i-- {
		fmt.Println(a[i])
	}
}
