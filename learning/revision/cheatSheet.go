package revision

import (
	"fmt"
	"time"
)

type Employee interface {
	getSalary()
}

type EmployeeDetails struct {
	name string
	age  int
}

func (s EmployeeDetails) getSalary() int {
	return 100
}

func hello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello times: ", i)
	}
}

func myChannelfunc(c chan string) {
	for i := range c {
		fmt.Println(i)
	}

}

func CheatSheet() {
	myStruct := EmployeeDetails{name: "Chetan", age: 30}
	EmployeeDetails.getSalary(myStruct)

	go hello()
	time.Sleep(1 * time.Second)

	mychan := make(chan string, 5)
	go myChannelfunc(mychan)
	for i := 0; i < 10; i++ {
		t := "Chetan times "
		time.Sleep(1 * time.Second)
		mychan <- t
	}
	close(mychan)

	a := []int{1, 2, 3, 4, 5, 6}
	for index, value := range a {
		fmt.Println(index, value)
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	b := map[string]string{}
	b["name"] = "Chetan"
	for k, v := range b {
		fmt.Println(k, v)
	}

	for k1, v1 := range b {
		if mval, boolCheck := b["name1"]; boolCheck {
			fmt.Println(k1, v1, mval)
			fmt.Println("Value is there")
		} else {
			fmt.Println("Value is NOT there")
		}
	}

}
