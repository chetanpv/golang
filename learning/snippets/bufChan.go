package snippets

import (
	"fmt"
	"time"
)

func BufChan() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	mychan := make(chan int, 4)
	go letsPrint(mychan)
	for _, value := range arr {
		mychan <- value
		time.Sleep(1 * time.Second) // This sleep is needed else the buffer becomes full and there will be deadlock
	}
	close(mychan)
}

func letsPrint(mychan chan int) {
	fmt.Println("Length of chan: ", len(mychan)) // 1
	for v := range mychan {
		fmt.Println(v) // 1 .. 8
	}
}
