package basics

import "fmt"

type friends struct {
	location string
}

type myStruct struct {
	name string
	age  int
	frnd friends
}

// Structures funtion shows how to use struct
func Structures() {
	details := myStruct{name: "Chetan", age: 31, frnd: friends{location: "Bangalore"}}
	fmt.Println(details)      // {Chetan 31 {Bangalore}}
	fmt.Println(details.name) // Chetan

	// Overriding default value
	details.name = "Chetan Patige"
	fmt.Println(details.name) // Chetan Patige
}
