package snippets

import (
	"fmt"
	"sort"
)

// Sorting ...
func Sorting() {
	a := []int{1, 6, 3, 4, 5, 1, 3, 5}
	sort.Ints(a)
	fmt.Println(a) //[1 1 3 3 4 5 5 6]
}
