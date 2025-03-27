package main

import "fmt"

const p = "len=%d cap=%d %v\n"

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf(p, len(s), cap(s), s)

	// o ':' é meio que um ponto de corte

	fmt.Printf(p, len(s[:0]), cap(s[:0]), s[:0])

	fmt.Printf(p, len(s[:4]), cap(s[:4]), s[:4])

	fmt.Printf(p, len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 110)
	fmt.Printf(p, len(s[:2]), cap(s[:2]), s[:2])
	fmt.Printf(p, len(s), cap(s), s)
}
