package main

import "fmt"

func somar(n1, n2 int) int {
	return n1 + n2
}

func calculadoraMultiplosRetornos(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}

func main() {

	numero := somar(10, 20)
	fmt.Println(numero)

	var f = func(a, b int) int {
		fmt.Println(somar(10, 20))
		return somar(a, b)
	}

	a := f(1, 2)
	fmt.Println(a)

	x, y, z, w := calculadoraMultiplosRetornos(10, 2)
	fmt.Println(x, y, z, w)

}
