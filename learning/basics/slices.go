package basics

import (
	"fmt"
)

// Slices ...
func Slices() {
	// --------------------------------------------- Initializing slice ---------------------------------------------
	s := make([]string, 3)
	// Also - s := []int{1,2,3}  // This is slice too
	fmt.Println("Emp: ", s) // [   ]

	// --------------------------------------- Adding values to slice via index---------------------------------------
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(s)      // [a b c]
	fmt.Println(s[2])   // c
	fmt.Println(len(s)) // 3

	// ---------------------------------- Dynamic adding values beyond initialization ----------------------------------
	s = append(s, "d")
	s = append(s, "e")
	fmt.Println(s) // [a b c d e]

	// -------------------------------------- Copying slices to another one --------------------------------------
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c) // cpy: [a b c d e]

	// -------------------------------------- Slicing the slice --------------------------------------
	fmt.Println(c[2:])  // [c d e]
	fmt.Println(c[1:3]) // [b c]
	fmt.Println(c[:3])  // [a b c]

	// -------------------------------------- Iterate over slice --------------------------------------
	for index, value := range c {
		fmt.Println(index, value)
		/*
			0 a
			1 b
			2 c
			3 d
			4 e
		*/
	}

	// -------------------------------------- Inserting a value in the middle --------------------------------------
	array1 := []int{1, 3, 4, 5}
	array2 := []int{2, 4, 6, 8}

	array1 = append(array1, 0)   // Step 1
	copy(array1[2:], array1[1:]) // Step 2
	array1[1] = array2[2]        // Step 3

	fmt.Println(array1) // [1 6 3 4 5]

	// -------------------------- Appending slice at the end to the original slice --------------------------
	d := []int{1, 2, 3}
	e := []int{4, 5, 6}
	for _, value := range e {
		d = append(d, value)
	}
	fmt.Println(d) // [ 1 2 3 4 5 6 ]

	// -------------------------- Prepend value to the slice -----------------------------
	f := []int{1, 2, 3, 4}
	f = append([]int{5}, f...)
	fmt.Println(f) // [5 1 2 3 4]
}
