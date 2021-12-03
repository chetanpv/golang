package basics

import "fmt"

// Arrays ...
func Arrays() {
	// ------------------------ Standard way to declare array and its default values ------------------------
	var a [5]int
	fmt.Println(a) // [0 0 0 0 0]
	var aa [3]string
	fmt.Println(aa) // [   ] // nil

	a[4] = 100
	fmt.Println(a) // [0 0 0 0 100]

	// ----------------------------- Initializing array with values -----------------------------
	b := [5]int{1, 2, 3, 4, 5} // This is array
	// example := []int{1, 2, 3, 4, 5} // This is slice
	fmt.Println(b) // [1 2 3 4 5]

	// ----------------------------- Ellipsis - Dynamic Array initialization with values -----------------------------
	c := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(c) // [ 1 2 3 4 5 6]

	// ----------------------------- Length of the array -----------------------------
	fmt.Println(len(c)) // 6

	// -------------------------------------- Common errors --------------------------------------

	// var e [...]int // use of [...] array outside of array literal
	// e[0] = 10

	//  -------------------------------------- Adding an element --------------------------------------
	// As arrays are already defined we can just edit index, c[2] = 200 , cannot actually append element

	// -------------------------------------- Removing an element in an array --------------------------------------
	// You need not remove element from array as such because space is already
	// laoccupied unlike slice. You can replace it if needed
	var e [3]string
	e[0] = "Chetan"
	e[1] = "Patige"
	fmt.Println(e) // [Chetan Patige ]
	e[1] = ""
	fmt.Println(e) // [Chetan   ]

	// -------------------------------------- Replacing an element --------------------------------------
	c[0] = 100
	fmt.Println(c) // [100 2 3 4 5 6]

	// -------------------------------------- Iterating over array --------------------------------------
	for index, value := range c {
		fmt.Println(index, value)
		/*
			0 100
			1 2
			2 3
			3 4
			4 5
			5 6
		*/
	}

	// Array comparison
	var f [2]int
	f[0] = 5
	var g [2]int
	g[0] = 5
	if f == g {
		fmt.Println(" Both have same array elements")
		fmt.Println("Address of f: ", f)
		fmt.Println("Address of g: ", g)
	}
}
