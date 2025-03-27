package main

// Usar ponteiro, quando quer mudar valor de algo (Mutável)
func somaComPonteiro(a, b *int) int {
	*a = 50
	*b = 40
	return *a + *b
}

// Usa uma cópia, não alterando os valores originais (Imutável)
func somaSemPonteiro(a, b int) int {
	return a + b
}

func main() {
	var1 := 10
	var2 := 20
	println(somaSemPonteiro(var1, var2))

	println(somaComPonteiro(&var1, &var2))

	println(var1)
	println(var2)

}
