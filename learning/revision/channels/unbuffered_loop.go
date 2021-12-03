package channels

import (
	"fmt"
	"time"
)

func UnbufferedLoop()  {
	mychan := make(chan string)
	go printChan(mychan)
	mychan <- "Chetan"
	mychan <- "Bindu"
	time.Sleep(1*time.Second)
	// Note
	// - If you send 2 strings into a channel and only one is listening, that is a deadlock as there is no one to listen
	// - Two print statements work fine. go routine makes a blocking call to listen

	mychan1 := make(chan string)
	go printChan1( mychan1)
	mychan1<- "Chetan1"
	mychan1<- "Chetan2"
	time.Sleep(1*time.Second)

	// Note:
	// - Range operator works fine with unbuffererd channels

}

func printChan1(mychan1 chan string) {
	for chanContent:= range mychan1 {
		fmt.Println("Chan1 content: ", chanContent)
	}
}

func printChan(mychan chan string) {
	fmt.Println("Channel contents: ", <-mychan)
	fmt.Println("Channel contents: ", <-mychan)
}
