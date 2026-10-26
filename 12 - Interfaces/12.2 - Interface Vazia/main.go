package main

import "fmt"

func main() {

	var (
		numero1 int
		numero2 any
	)

	numero1 = numero1 + 1
	numero1++

	// numero2++ não é tratado como numero
	numero2Int := numero2.(int)
	numero2Int++

	// --------

	var (
		slice1 any = []int{1, 2, 3}
		slice2 any = []int{4, 5, 6}
	)

	// isso não pode
	if slice1 == slice2 {
		fmt.Println("Os slices são iguais")
	} else {
		fmt.Println("Os slices são diferentes")
	}

	var x any
	x = numero1
	x = "oi"
	x = true

	switch v := x.(type) {
	case int:
		fmt.Println("O valor é um inteiro:", v)
	case string:
		fmt.Println("O valor é uma string:", v)
	case bool:
		fmt.Println("O valor é um booleano:", v)
	default:
		fmt.Println("O valor é de outro tipo")

	}
}
