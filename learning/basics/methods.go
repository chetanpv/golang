package basics

import "fmt"

type rect struct {
	length  int
	breadth int
}

type square struct {
	length int
	breadth int
}

// when we put (r rect) , the function area gets tied up with the structure rect. This function is of type rect struct
// Anyone creating a structure variable like rectangle here, can access these functions.
func (r rect) area() {
	fmt.Println(r.length * r.breadth)
}

func (r rect) perimeter() {
	fmt.Println(2*r.length + 2*r.breadth)
}

func (r square) area() {
	fmt.Println(r.length * r.breadth)
}

func (r square) perimeter() {
	fmt.Println(2*r.length + 2*r.breadth)
}

// Methods shows how methods are differnt from functions
func Methods() {
	rectangle := rect{length: 2, breadth: 3}
	fmt.Println(rectangle.length)

	rectangle.area()
	rectangle.perimeter()

}
