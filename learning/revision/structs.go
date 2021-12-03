package revision

import "fmt"

type Animal struct {
	leg int
}

type Geometry struct {
	sides *int
}

func Structs() {
	dog := Animal{leg: 4}
	fmt.Println(dog)

	squareSides := 4
	square := Geometry{sides: &squareSides}
	fmt.Println(*square.sides)
	newSquareSides := 3
	square.sides = &newSquareSides
	fmt.Println(*square.sides)
}
