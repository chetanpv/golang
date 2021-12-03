package channels

import (
	"fmt"
	"time"
)

func Select() {
	mychan:= make(chan string)
	mychan1:= make(chan string)
	go printOne(mychan)
	go printTwo(mychan1)
	select {
	// This is a receiver from the channel
	case one := <- mychan:
		fmt.Println(one)

	case two:= <- mychan1:
		fmt.Println(two)
	//default:
	//	fmt.Println("No channels to serve")
	}
	// Note:
	// - The select statement can be used to send data over channels on which go routine is free.
	// - Only one condition need to be true to execute.
	// - Can be used in queuing mechanism
	// - In our example both the go routines were executed and got blocked, which ever channel listened first served the object
		// though both go routines are ready to listen
}

func printTwo(mychan1 chan string) {
	time.Sleep(1* time.Second)
	mychan1 <- "This is from print two"
}

func printOne(mychan chan string) {
	time.Sleep(2* time.Second)
	mychan <- "This is from print one"
}


