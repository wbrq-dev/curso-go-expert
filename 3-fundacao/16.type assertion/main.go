package main

import "fmt"

func main() {
	var minhaVar interface{} = "Williams Barriquero"
	println(minhaVar.(string))

	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %d e o resultado de ok é %t\n", res, ok)

	res2 := minhaVar.(int)
	fmt.Printf("O valor de res2 é %v", res2)
}
