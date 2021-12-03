package revision

import "fmt"

type geometry interface {
	area()
	perimeter()
}

type details struct {
	length  int
	breadth int
}

func (d details) area() {
	fmt.Println(d.length * d.breadth)
}

func (d details) perimeter() {
	fmt.Println(d.length + d.breadth)
}

func Interfaces() {
	a := details{length: 2, breadth: 3}
	geometry.area(a)

	b := details{length: 5, breadth: 4}
	geometry.perimeter(b)
}
