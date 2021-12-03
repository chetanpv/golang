package channels

import (
	"fmt"
	"time"
)

func MulGoRoutines() {
	mychan := make(chan string)
	go ReadOne(mychan)
	go ReadTwo(mychan)
	mychan <- "Chetan"
	mychan <- "Bindu"
	mychan <- "Chetan1"
	mychan <- "Bindu1"
	time.Sleep(1*time.Second)

	//From Read Two:  Chetan
	//From Read Two:  Chetan1
	//From Read One:  Bindu
	//From Read Two:  Bindu1

	// Note:
	// As you can see, the goroutine can be read from any channel
}

func ReadOne(mychan chan string) {
	for element:= range mychan {
		fmt.Println("From Read One: ", element)
	}
}

func ReadTwo(mychan chan string) {
	for element:= range mychan {
		fmt.Println("From Read Two: ", element)
	}
}