package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Desafio
// Função que busque um item em uma lista e retorne sua posição
// [10,20,30]-> 20 posicao -> 1
// Qualquer tipo inteiro,strings, Pessoa c nome e idade
type (
	Pessoa struct {
		Nome  string
		Idade int
	}

	cInt constraints.Integer
)

func leCoisas[T cInt | Pessoa | string](coisas []T, coisa T) int {
	for i, v := range coisas {
		if v == coisa {
			return i
		}
	}
	return -1
}

func main() {

	var (
		listaPessoa = []Pessoa{
			Pessoa{
				Nome:  "Guilherme",
				Idade: 5,
			},
			Pessoa{
				Nome:  "Maria",
				Idade: 6,
			},
		}

		listaNumeros = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		listaString  = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	)

	fmt.Println(leCoisas(listaPessoa, Pessoa{Nome: "Guilherme", Idade: 5}))
	fmt.Println(leCoisas(listaNumeros, int64(5)))
	fmt.Println(leCoisas(listaString, "5"))

}

/*

type (
	Pessoa struct {
		Nome  string
		Idade int
	}

	Buscavel interface{
		constraints.Integer | ~string | Pessoa
)

func Busca[T Buscavel](lista[]T, f func(T) bool) int{
	for i, valor := range lista{
		if f(valor){
			return i
		}
	}
	return -1
}

*/
