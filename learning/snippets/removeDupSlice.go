package snippets

import "fmt"

// Remove all the duplicates from the slice and give end result as a set

// RemoveDup...
func RemoveDup() {
	myint := []int{1, 2, 3, 1, 2, 3, 4}
	mylist := removeDuplicateInt(myint)
	fmt.Println(mylist)
}

func removeDuplicateInt(intSlice []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
