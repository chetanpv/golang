package channels

import (
	"fmt"
	"sync"
	"time"
)

func Waitgroups() {
	var wg sync.WaitGroup

	for i:=0; i< 3; i++ {
		wg.Add(1)
		go tryservice1(&wg, i)
	}

	wg.Wait()
	fmt.Println("Stopped")
}

func tryservice1(s *sync.WaitGroup, i int) {
	time.Sleep(1* time.Second)
	fmt.Println("The attempt of calling this function: ", i)
	s.Done()
}
