package snippets

import "fmt"

func NumberOcc() {
	number := []int{1, 2, 3, 4, 5, 2, 3, 4, 7, 8, 5}
	result := map[int]int{}

	for _, value := range number {
		if mapValue, mapBoolCheck := result[value]; mapBoolCheck {
			result[value] = mapValue + 1
		} else {
			result[value] = 1
		}
	}

	fmt.Println(result)
}

func CharOcc() {
	stringList := []string{"a", "b", "a", "c", "b", "b"}
	result := map[string]int{}

	for _, value := range stringList {
		if mapValue, mapBoolCheck := result[value]; mapBoolCheck {
			result[value] = mapValue + 1
		} else {
			result[value] = 1
		}
	}

	fmt.Println(result)

	mystring := "ChetanChetan"
	myresult := map[rune]int{}
	for _, mystringValue := range mystring {
		if mymapValue, mymapBoolCheck := myresult[mystringValue]; mymapBoolCheck {
			myresult[mystringValue] = mymapValue + 1
		} else {
			myresult[mystringValue] = 1
		}
	}

	newResult := map[string]int{}
	for key, keyValue := range myresult {
		newResult[string(key)] = keyValue
	}
	fmt.Println(newResult)
}
