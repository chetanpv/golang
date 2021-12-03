package channels

import (
	"fmt"
	"time"
)

func Unbuffered() {
	mychan := make(chan string)
	go printString(mychan)
	mychan <- "Chetan"
	time.Sleep(1* time.Second) // If this statement is not there, the program will end without error
	// Note:
	// - something can be pushed into a channel only when there is a listener as unbuffered channels do a blocking calls
			// In this case read was blocked as there was no listener
	// - Give time for the listener to process what was sent, ending the main program before that will end channel
}

func printString(mychan chan string) {
	fmt.Println("This is my channel string: ", <- mychan)
}
