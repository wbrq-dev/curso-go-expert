package main

import "fmt"

type Endereco struct {
	Logradouro string
	Cidade     string
	Estado     string
	Numero     int
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco // Compondo
	//Address  Endereco // Atributo do tipo Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Println(c)
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

func (e Empresa) Desativar() {
	fmt.Println("Método não implementado")
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {

	williams := Cliente{
		Nome:  "Williams",
		Idade: 29,
		Ativo: true,
	}

	minhaEmpresa := Empresa{
		Nome: "RPE",
	}

	//fmt.Println(williams)
	//
	//williams.Cidade = "Camaragibe"
	//williams.Desativar()
	//
	//fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", williams.Nome, williams.Idade, williams.Ativo)
	Desativacao(williams)

	Desativacao(minhaEmpresa)
}
