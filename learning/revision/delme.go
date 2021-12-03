package revision

import (
	"fmt"
)

func Delme() {
	mins := make(chan int)
	maxs := make(chan int)
	done := make(chan int)
	go produce(mins, done, []int{2, 5, 1, 4})
	go produce(maxs, done, []int{4, 8, 9, 7})
	for n := 2; n > 0; {
		select {
		case p := <-mins:
			fmt.Println("Min:", p)	//consume output
		case p := <-maxs:
			fmt.Println("Max:", p)	//consume output
		case <-done:
			n--
		}
	}
}

func produce(ch, done chan int, data []int) {
	for _, v := range data {
		ch <- v
	}
	done <- 1
}
