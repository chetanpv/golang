package snippets

import (
	"fmt"
	"time"
)

var flag bool = true

func Singleton() {
	go mysingleton("Chetan")
	go mysingleton("Chetan1")
	time.Sleep(100 * time.Second)
}

func mysingleton(mymethod string) {
	for i := 0; i < 4; i++ {
		if flag == true {
			fmt.Println("I am free for ", mymethod)
			fmt.Println("Going to implement something. Locking the flag")
			flag = false
			time.Sleep(5 * time.Second)
			flag = true
		} else if flag == false {
			fmt.Println("Something is running. Wait !!!")
		}
	}

}
