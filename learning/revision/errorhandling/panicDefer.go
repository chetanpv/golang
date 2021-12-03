package errorhandling

import "fmt"

func PanicDefer() {
	defer letsCreatePanic()
	panic ("This will be printed last")

	// Note :
	// - Panic will always be last statement
	// - Defer will be last but one when panic is present

}

func letsCreatePanic() {
	fmt.Println("This will print last but one")
}

//This will print last but one
//panic: This will be printed last
//
//goroutine 1 [running]:
//learning/revision/errorhandling.PanicDefer()
//C:/Chetan/Work/Go/src/learning/revision/errorhandling/panicDefer.go:7 +0x4a
//main.main()
//C:/Chetan/Work/Go/src/learning/main.go:60 +0x17
//
//Process finished with the exit code 2
