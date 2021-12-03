package snippets

import "fmt"

func ChanOddEven() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mychannel := make(chan int)
	go evenOdd(mychannel)
	for _, value := range a {
		mychannel <- value
	}
	close(mychannel)
}

func evenOdd(mychan chan int) {
	// Range over channels will be there till we close the channel !! That is why sleep was not needed in this program
	for v := range mychan {
		if v%2 == 0 {
			fmt.Println("Number", v, "is even")
		} else {
			fmt.Println("Number", v, "is odd")
		}
	}
}

// func evenOdd(mychan chan int, length int) {
// 	for i := 0; i <= length; i++ {
// 		value := <-mychan
// 		fmt.Println("Value: ", value)
// 		if value%2 == 0 {
// 			fmt.Println("Number", value, "is even")
// 		} else {
// 			fmt.Println("Number", value, "is odd")
// 		}
// 	}
// }
