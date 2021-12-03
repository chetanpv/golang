package errorhandling

import "fmt"

func RecoverFromPanic() {
	defer PanicHandler()
	panic("We are heading towards panic")
}

func PanicHandler() {
	if err:= recover() ; err != nil {
		fmt.Println("We are saved!! The culprit was:\n" , err)
	}
}
