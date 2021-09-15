package basics

import "fmt"

// Loop funtions shows differnt ways to loop. There is no while loop in golang
func Loop() {
	// C language like loop
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			fmt.Println(n)
		}
	}

	// while loop kind. No while loop in golang
	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}

	// Infinite loop
	for {
		fmt.Println("Hello")
		break
	}
}
