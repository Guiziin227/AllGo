package main

import "fmt"

//Go não tem herança, mas tem composição

type pessoa struct {
	Nome  string
	Idade int
}

type estudante struct {
	pessoa
	Curso string
}

func main() {

	p := pessoa{Nome: "Malu",
		Idade: 20}
	e := estudante{
		pessoa: p,
		Curso:  "Arq",
	}

	fmt.Println(e.Curso)
	fmt.Println(e.Nome)
	fmt.Println(e.Idade)
	fmt.Println(e)
}
