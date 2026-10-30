package main

import "fmt"

func SomaInteiros(numeros []int) int {
	var soma = 0
	for _, n := range numeros {
		soma += n
	}
	return soma
}

func SomaGenerics[T int | float64](numeros []T) T {
	var soma T = 0
	for _, n := range numeros {
		soma += n
	}
	return soma
}

func main() {

	numerosInteiros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numerosFloat := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}

	soma := SomaInteiros(numerosInteiros)
	fmt.Printf("Soma: %v\n", soma)

	fmt.Println(SomaGenerics(numerosInteiros))
	fmt.Println(SomaGenerics(numerosFloat))

}
