package snippets

import "fmt"

func ReverseArr() {
	a := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(a)
	for i := 0; i < len(a)/2; i++ {
		j := len(a) - i - 1
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)
}
