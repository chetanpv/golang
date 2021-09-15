package snippets

import "fmt"

// AddRemoveSlice ...
func AddRemoveSlice() {
	// --------------------------- Append at the end ---------------------------
	a := []int{1, 2, 3, 4}
	b := []int{5, 6, 7}
	for _, value := range b {
		a = append(a, value)
	}
	fmt.Println(a) // [1 2 3 4 5 6 7]

}
