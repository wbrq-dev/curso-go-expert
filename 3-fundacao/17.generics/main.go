package main

type MyNmber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma

}

func Compara[T comparable](a, b T) bool {
	return a == b
}

func main() {
	m := map[string]int{"Williams": 100, "Barriquelo": 200, "Gomes": 300}
	m2 := map[string]float64{"Williams": 100, "Barriquelo": 200, "Gomes": 300}

	m3 := map[string]MyNmber{"Williams": 100, "Barriquelo": 200, "Gomes": 300}
	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10))
}
