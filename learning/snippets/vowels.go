package snippets

import "fmt"

func Vowels() {
	mystring := "Chetan"
	result := map[rune]int{'a': 0, 'e': 0, 'i': 0, 'o': 0, 'u': 0}
	for _, v := range mystring {
		// Note: it has to be 'a' and not "a". The type is rune and not string
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			result[v] = result[v] + 1
		}
	}

	finalResult := map[string]int{}
	for key, val := range result {
		finalResult[string(key)] = val
	}
	fmt.Println(finalResult)
}
