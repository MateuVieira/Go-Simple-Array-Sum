package main

import (
	"fmt"
)

/*
 * Complete the simpleArraySum function below.
 */
func simpleArraySum(ar []int) int {
	var result int = 0

	for _, number := range ar {
		result += number
	}

	return result
}

func main() {
	ar := []int{1, 2, 3, 4, 10, 11}

	result := simpleArraySum(ar)

	fmt.Printf("%d\n", result)
}
