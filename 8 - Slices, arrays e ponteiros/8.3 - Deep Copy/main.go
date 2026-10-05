package main

import "fmt"

func main() {

	matriz := [][]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println(len(matriz))
	fmt.Println(len(matriz[0]))

	matriz2 := make([][]int, len(matriz))

	matriz2 = deepCopy(matriz)

	matriz2[0][0] = 100

	fmt.Println(matriz, matriz2)
}

func deepCopy(matriz [][]int) [][]int {
	destino := make([][]int, len(matriz))
	for i, row := range matriz {
		destino[i] = make([]int, len(row))
		copy(destino[i], row)
	}
	return destino
}
