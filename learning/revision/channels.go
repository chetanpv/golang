package revision

import (
	"fmt"
	"time"
)

func Channel() {
	chan1 := make(chan string)
	go try1(chan1)
	chan1 <- "Chetan"
	chan1 <- "Chetan1"
	chan1 <- "Chetan2"
	close(chan1)

	chan2 := make(chan string, 5)
	chan2 <- "Bindu"
	chan2 <- "Bindu1"
	chan2 <- "Bindu2"
	chan2 <- "Bindu3"
	chan2 <- "Bindu4"
	fmt.Println("Length of chan2: ", len(chan2))
	go try2(chan2) // Unlike unbuffered channels, we can send data even if listener is not there. Max 5 in this case
	time.Sleep(2 * time.Second)
	chan2 <- "Final"
	go try3(chan2)
	time.Sleep(2 * time.Second)
}

func try1(chan1 chan string) {
	fmt.Println("Chan1: ", <-chan1)
	fmt.Println("Chan1: ", <-chan1)
	fmt.Println("Chan1: ", <-chan1)
}

func try2(chan2 chan string) {
	for i := range chan2 {
		fmt.Println("Chan2: ", i)
	}
}

func try3(chan3 chan string) {
	fmt.Println("Final: ", <-chan3)
}
