package channels

import (
	"fmt"
	"time"
)

func Timeouts() {
	chan1 := make(chan string)
	chan2 := make(chan string)
	go service1(chan1)
	go service2(chan2)

	select {
	case one:= <- chan1:
		time.Sleep(1*time.Second)
		fmt.Println(one)
	case two:= <- chan2:
		time.Sleep(2*time.Second)
		fmt.Println(two)
	case <- time.After(3*time.Second):
		fmt.Println("There are no go routines available")
	}
	// Note :
	// THe above cases were blocked till case 3 which opened at 3 seconds.. that go executed
	// This can be used as timeout
}

func service2(chan2 chan string) {
	time.Sleep(4*time.Second)
	chan2 <- "From service 2"
}

func service1(chan1 chan string) {
	time.Sleep(4*time.Second)
	chan1 <- "From service 1"
}
