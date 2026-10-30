package main

import "fmt"

type myint int

type Number interface {
	~int | float64
}

func Soma[T Number](a, b T) T {
	return a + b
}

func main() {

	var (
		x1 = int(100)
		x2 = int(200)

		y1 = float64(300.4)
		y2 = float64(400.5)

		z1 = myint(100)
		z2 = myint(200)
	)

	fmt.Println(Soma(x1, x2))
	fmt.Println(Soma(y1, y2))
	// fmt.Println(Soma(x1, y1)) não pode.

	fmt.Println(Soma(z1, z2))
}
