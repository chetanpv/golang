package channels

import (
	"fmt"
)

func FanIntoOneChannel() {
	chan1 := make(chan string)
	counter := 0
	aVal,bVal,cVal,dVal,eVal := 0, 0, 0, 0,0

	go Tryme(chan1)
	for  {
		select {
		case <- chan1:
			aVal ++
		case <- chan1:
			bVal ++
		case <- chan1:
			cVal ++
		case <- chan1:
			dVal ++
		case <- chan1:
			eVal ++
		}
		counter += 1
		if counter == 100 {
			fmt.Println("Total times run: ", counter)
			fmt.Println("Values: a: ",  aVal, "b: ", bVal, "c: ", cVal, "d: ", dVal, "e: ", eVal)
			break
		}

	}

}

func Tryme(chan1 chan string) {
	for i:= 0; i< 100 ; i++ {
		chan1 <- "Trying to process"
	}
}

/*
Notes :
- This program took 100 calls to same goroutine(can be different, does not matter) and processed into one channel
- This is simple fan in concept
- Used in Queues
*/

