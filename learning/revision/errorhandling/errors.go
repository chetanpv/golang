package errorhandling

import (
	"errors"
	"fmt"
)

func ErrorTest() {
	err := printError()
	if err != nil {
		fmt.Println("The error is: ", err)
	}
	// Note:
	// - This is a sample way to create a new error and print it
}

func printError() error {
	err := errors.New("This is an error")
	return err
}
