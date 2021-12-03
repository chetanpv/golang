package basics

import "fmt"

// Channels ...
func Channels() {
	// ------------------------------------ Simple channel example ------------------------------------
	mych1 := make(chan string)
	// Function needs to be called first else you are sending something to a channel without receiver
	go tryOne(mych1)

	// Note: We already called the go routine without sending any data
	// The channel is a blocker call. So, the go routine does not exit
	mych1 <- "AttemptOne"
	mych1 <- "AttemptTwo"
	//mych1 <- "AttemptThree" // This will cause deadlock

	// Channels need to be closed manually.
	// If not the channel assumes that some one will send data to it
	close(mych1)

	// ------------------------------------ Looping over channels ------------------------------------
	mych2 := make(chan string)
	go tryTwo(mych2)
	arr := []string{"Hello", "world", "how", "are", "you"}
	for _, value := range arr {
		mych2 <- value
	}
	close(mych2)

	// ------------------------------------ Buffered channels ------------------------------------
	mych3 := make(chan int, 3)
	go tryThree(mych3)
	mych3 <- 1
	mych3 <- 2
	mych3 <- 3
	// If you execute program at this line, nothing will be printed in the function.
	// Channel flushes out only when the pipe is full. It is like channel needs n+1 to execute
	mych3 <- 4
	mych3 <- 5
	mych3 <- 6
	// After the above line, 1 2 3 gets printed.. Strange though.. :(
	close(mych3)
}

// Reading value from a channel
func tryOne(c chan string) {
	fmt.Println("Data from tryOne: ", <-c)
	fmt.Println("Data from tryOne: ", <-c)
}

// Reading value from a channel in a loop
func tryTwo(c chan string) {
	// channel for 5 values
	for i := 1; i <= 5; i++ {
		fmt.Println("Data from tryTwo: ", <-c)
	}
}

func tryThree(c chan int) {
	for i := 0; i <= 3; i++ {
		fmt.Println("Data from tryThree: ", <-c)
	}
}
