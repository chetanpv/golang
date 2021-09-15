package snippets

import "fmt"

func Fibonacci(number int) {
	result := []int{1, 2}
	for i := 1; i <= number; i++ {
		num := result[len(result)-1] + result[len(result)-2]
		result = append(result, num)
	}
	fmt.Println(result)
}
