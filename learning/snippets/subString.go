package snippets

import (
	"fmt"
	"strings"
)

func SubString() {
	mystring := "ChetanPatige"
	mysubstring := "tan"

	if strings.Contains(mystring, mysubstring) {
		fmt.Println("Contains")
	} else {
		fmt.Println("Does not Contain")
	}
}
