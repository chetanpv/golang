package basics

import (
	"errors"
	"fmt"
)

// ErrorHandling ...
func ErrorHandling() {
	errorFunction()
	panicDeferRecoverFunction()
}

func errorFunction() {
	err := validateError()
	if err != nil {
		fmt.Println(err)
	}
}

func validateError() error {
	if 5 < 6 {
		err := errors.New("This is an error. Not really")
		return err
	}
	return nil
}

func panicDeferRecoverFunction() {
	// Defer
	/*
		A defer statement is a mechanism used to defer a function by putting it into an executed stack once
		the function that contains the defer statement has finished, either normally by executing a return
		statement or abnormally panicking. Deferred functions will then be executed in reverse order in which they were deferred.
	*/
	defer recoveringPanic()

	// Panic
	/*
		The panic function will be called by the program if code has run time issues
		OR user can define it
	*/
	panic("Not enough statements. The code will stop here")
}

func recoveringPanic() {
	// Recover
	/*
		User can create recover functions/statements if the code should not stop from panic and handle exceptions accordingly
	*/

	if a := recover(); a != nil {
		fmt.Println("RECOVER: ", a) // RECOVER:  Not enough statements. The code will stop here
	}
}
