package snippets

import "fmt"

func Sets() {
	myarr := []int{1, 2, 3, 4, 5, 1}
	tempMap := map[int]bool{}
	mySet := []int{}
	for _, value := range myarr {
		if _, mapBool := tempMap[value]; !mapBool {
			tempMap[value] = true
			mySet = append(mySet, value)
		}
	}
	fmt.Println(mySet)

}
