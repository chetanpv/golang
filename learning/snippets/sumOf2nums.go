package snippets

import "fmt"

func SumOfTwoNums() {
	myarr := []int{-1, 0, 5, 1, 2, 6, 7}
	k := 6

	for index, value := range myarr {

		indexValue := value
		mytempArr := myarr[index:]

		for _, value1 := range mytempArr {
			if value1+value == k {
				fmt.Println(indexValue, " and ", value1, "=", k)
			}
		}
	}
}
