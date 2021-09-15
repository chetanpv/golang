package basics

import "fmt"

// Maps ...
func Maps() {
	myMap1 := map[string]string{}
	myMap3 := make(map[string]string)
	myMap4 := map[string]string{"name": "Chetan", "middleName": "Patige"}

	// ----------------------------------- Adding value to a map -----------------------------------
	myMap1["name"] = "Chetan"
	fmt.Println(myMap1) //map[name:Chetan]

	// ----------------------------------- Delete a value from a map -----------------------------------
	myMap3["name"] = "Chetan"
	delete(myMap3, "name")
	fmt.Println(myMap3) // map[]

	// ----------------------------------- Iterate over the maps -----------------------------------
	for index, value := range myMap4 {
		fmt.Println(index, value)
	}
}
