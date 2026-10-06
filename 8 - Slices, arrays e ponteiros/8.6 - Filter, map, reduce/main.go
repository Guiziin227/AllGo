package main

import "fmt"

type meuSlice []int

func main() {

	var lista = meuSlice{1, 2, 3, 4, 5}

	fmt.Println(lista)

	lista = lista.Filter(func(x int) bool {
		return x%2 == 0
	})

	fmt.Println(lista)

	lista = lista.Map(func(x int) int {
		return x * 2
	})

	fmt.Println(lista)

	soma := lista.Reduce(func(x, y int) int {
		return x + y
	}, 0)

	fmt.Println(soma)

	multiplicacao := lista.Reduce(func(x, y int) int {
		return x * y
	}, 1)

	fmt.Println(multiplicacao)

	fmt.Println("==========================")

	var lista2 = meuSlice{1, 2, 3, 4, 5}

	mult2 := lista2.Filter(func(x int) bool {
		return x%2 == 0
	}).Map(func(x int) int {
		return x * 2
	}).Reduce(func(i int, i2 int) int {
		return i * i2
	}, 1)

	fmt.Println(mult2)
}

func (s meuSlice) Filter(cond func(int) bool) meuSlice {
	var result meuSlice

	for _, valor := range s {
		if cond(valor) {
			result = append(result, valor)
		}
	}

	return result
}

func (s meuSlice) Map(trans func(int) int) meuSlice {
	var result meuSlice

	for _, val := range s {
		result = append(result, trans(val))
	}

	return result
}

func (s meuSlice) Reduce(acc func(int, int) int, inicial int) int {
	var result = inicial

	for _, val := range s {
		result = acc(result, val)
	}

	return result
}
