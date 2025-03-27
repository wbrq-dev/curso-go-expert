package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2, 3, 45, 567, 67, 88, 9, 5, 6, 34, 52, 3, 4213, 421, 34))
}

func sum(numeros ...int) int {
	total := 0

	for _, n := range numeros {
		total += n
	}

	return total
}
