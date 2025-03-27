package main

import "fmt"

func main() {
	salarios := map[string]int{"Williams": 14500, "Viviane": 14455}
	fmt.Println(salarios["Williams"])
	//delete(salarios, "Williams")
	salarios["Will"] = 14567
	fmt.Println(salarios["Will"])

	sal := make(map[string]int)
	sal1 := map[string]int{}
	sal1["Williams"] = 14500
	sal["Williams"] = 14587

	fmt.Println(sal["Williams"])
	fmt.Println(sal1["Williams"])

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salário é de %d\n", salario)
	}
}
