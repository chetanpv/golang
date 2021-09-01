package basics

import "fmt"

type rect1 struct {
	length  int
	breadth int
}

type myRect interface {
	area1()
	perimeter1()
}

// when we put (r rect) , the function area gets tied up with the structure rect. This function is of type rect struct
// Anyone creating a structure variable like rectangle here.. can access these functions.
func (r rect1) area1() {
	fmt.Println(r.length * r.breadth)
}

func (r rect1) perimeter1() {
	fmt.Println(2*r.length + 2*r.breadth)
}

// Interfaces function shows how to use a simple interface
func Interfaces() {
	rectangle1 := rect1{length: 2, breadth: 3}
	myRect.area1(rectangle1)
}
