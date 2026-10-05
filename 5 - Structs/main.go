package main

import "fmt"

type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereco
}

type Endereco struct {
	Rua    string
	Bairro string
}

func main() {
	p := Pessoa{Nome: "Gui", Idade: 20, Endereco: Endereco{Rua: "Rua 1", Bairro: "Bairro 1"}}
	fmt.Println(p.Endereco.Rua)
	fmt.Println(p.Endereco.Bairro)
}
