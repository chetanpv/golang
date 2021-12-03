package channels

import (
	"fmt"
	"time"
)

func Buffered() {
	mychan := make(chan string, 3)
	go printMyChan(mychan)
	mychan <- "Chetan"
	time.Sleep(1*time.Second)
	close(mychan)

	//	Note:
	// - You need not wait for the channel buffer to fill. Message is popped out if there is a listener

	mychan1 := make(chan string, 3)
	//go printMyChan1(mychan1) // This is not required as we are sending only the beffered size into channel and not reading it
	mychan1 <- "Chetan1"
	mychan1 <- "Chetan2"
	mychan1 <- "Chetan3"
	close(mychan1)
	// Note:
	// - This wont print anything because the goroutines are not blocked till the buffer is complete
	// - Error wont be thrown even if the main exits
	// - Reason : Buffered channels are non blocking till the buffer is full+1

	// Deadlock:
	// If you are trying to read data from a channel but channel does not have a value available with it,
	// it blocks the current goroutine and unblocks other in a hope that some goroutine will push a value to the channel.
	// Hence, this read operation will be blocking. Similarly, if you are to send data to a channel,
	// it will block current goroutine and unblock others until some goroutine reads the data from it.
	// Hence, this send operation will be blocking.

	mychan2 := make(chan string, 3)
	go printMyChan2(mychan2) // If I comment this, there will be deadlock because goroutine will be locked at Chetan44
	mychan2 <- "Chetan11"
	mychan2 <- "Chetan22"
	mychan2 <- "Chetan33"
	mychan2 <- "Chetan44"
	time.Sleep(1*time.Second)
	close(mychan2)
}

func printMyChan2(mychan2 chan string) {
	for val :=range mychan2{
		fmt.Println("Mychan1 value: ", val)
	}
}

func printMyChan1(mychan1 chan string) {
	for val :=range mychan1{
		fmt.Println("Mychan1 value: ", val)
	}
}

func printMyChan(mychan chan string) {
	fmt.Println("MyChan content: ", <-mychan)
}
