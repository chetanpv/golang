package channels

import (
	"fmt"
	"time"
)

func UniDirectional() {
	mychan1:= make(chan string)
	go readingChan(mychan1)
	mychan1 <- "Chetan"
	time.Sleep(1 * time.Second)

}

func readingChan(mychan1 <- chan string) {
	fmt.Println("This can only read: ", <- mychan1)
}
