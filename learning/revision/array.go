package revision

import "fmt"

func Array() {
	a := []int{}
	//a[0] = 1 // This is wrong. As we have initialized slice and indexing on 0
	a = append(a, 1)
	fmt.Println(a) // [1]

	b := []int{1, 2, 3, 4, 5}
	fmt.Println(b) // [1 2 3 4 5]

	var c [2]int
	c[0] = 1
	fmt.Println(c) // [1 0]

	// Iter over an array
	d := []int{1, 2, 3, 4, 5}
	for index, value := range d {
		fmt.Println(index, value)
	}

	// Slicing an array
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(e[:])   // Everything
	fmt.Println(e[3:])  // from next
	fmt.Println(e[:5])  // till given number
	fmt.Println(e[3:6]) // from next to given number

}
