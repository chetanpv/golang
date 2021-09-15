package basics

import "fmt"

// Pointers function gives a simple example on how to use pointer
func Pointers() {
	name := "Chetan"
	fmt.Println(name)  // Chetan
	fmt.Println(&name) // 0xc0000361f0

	sendAddress(&name) // Send the address as the function accepts pointer
	fmt.Println(name)  // P V
}

/*
	Creating a function named sendAddress which takes a variable with name myname.
	The type of the variable is string.. and the kind is pointer.. hence *string
*/
func sendAddress(myname *string) {
	// myname is address. If we want to update it. we need to use pointer to udpate it.
	// Hence *myname, which goes to the address and updates the value
	*myname = "P V"
}
