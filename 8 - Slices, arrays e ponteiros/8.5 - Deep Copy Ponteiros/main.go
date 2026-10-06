package main

import "fmt"

type Pessoa struct {
	Nome  *string
	Idade int
}

func main() {

	var nome string = "Gui"

	pessoa1 := Pessoa{
		Nome:  &nome,
		Idade: 20,
	}

	pessoa2 := pessoa1

	fmt.Println(*pessoa1.Nome, pessoa1.Idade)
	fmt.Println(*pessoa2.Nome, pessoa2.Idade)

	pessoa2.Nome = toPointer("Malu")

	fmt.Println(*pessoa1.Nome, pessoa1.Idade)
	fmt.Println(*pessoa2.Nome, pessoa2.Idade)

	pessoa3 := Pessoa{
		Nome:  toPointer("Pedro"),
		Idade: 15,
	}

	pessoa4 := deepCopy(pessoa3)

	fmt.Println(*pessoa3.Nome, pessoa3.Idade)
	fmt.Println(*pessoa4.Nome, pessoa4.Idade)

	pessoa4.Nome = toPointer("Leandro")
	fmt.Println(*pessoa3.Nome, pessoa3.Idade)
	fmt.Println(*pessoa4.Nome, pessoa4.Idade)

	fmt.Println("======================")
	listaDePessoas := []Pessoa{pessoa1, pessoa2, pessoa3, pessoa4}

	listaDePessoas2 := deepCopyList(listaDePessoas)

	fmt.Println("Lista 1:", listaDePessoas)
	fmt.Println("Lista 2:", listaDePessoas2)

}

// toPointer nas novas versões de GO pode ser trocada por new()
func toPointer(s string) *string {
	return &s
}

func deepCopy(origem Pessoa) Pessoa {

	var destino Pessoa

	destino.Idade = origem.Idade

	destino.Nome = new(*origem.Nome)

	return destino
}

func deepCopyList(origem []Pessoa) []Pessoa {
	var destino = make([]Pessoa, len(origem))

	for i, p := range origem {
		destino[i] = deepCopy(p)
	}

	return destino
}
