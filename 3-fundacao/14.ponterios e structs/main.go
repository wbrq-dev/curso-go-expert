package main

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	return c.saldo
}

//func (c Conta) simularSemPonteiro(valor int) int {
//	c.saldo += valor
//	return c.saldo
//}

func main() {
	c := Conta{saldo: 100}
	println(c.simular(200))
	//println(c.simularSemPonteiro(200))
	println(c.saldo)
}
