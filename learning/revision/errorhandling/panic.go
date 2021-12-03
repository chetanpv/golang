package errorhandling

import "fmt"

func PanicCreate()  {
	panic("THis will end the program")
	fmt.Println("This wont get executed")
}
