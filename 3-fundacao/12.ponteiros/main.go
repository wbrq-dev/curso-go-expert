package main

import "fmt"

func main() {
	a := 10
	fmt.Println(a)
	var ponteiro = &a
	println(*ponteiro)

	*ponteiro = 20
	b := &a
	fmt.Println(*b)
	fmt.Println(a)
}
