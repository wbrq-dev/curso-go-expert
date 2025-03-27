package main

import "fmt"

func main() {
	var array [3]int
	array[0] = 10
	array[1] = 20
	array[2] = 10

	//len pega o tamanho do array
	fmt.Println(array[len(array)-1])

	for i, v := range array {
		fmt.Printf("O valor do indice é %d e o valor é %d\n", i, v)
	}
}
