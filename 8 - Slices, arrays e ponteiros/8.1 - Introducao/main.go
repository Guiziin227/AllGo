package main

import "fmt"

func main() {

	slice := []int{100, 200, 300, 400, 500}
	array := [5]int{100, 200, 300, 400, 500}

	fmt.Println("Tamanho do array =", len(array))
	fmt.Println("Capacidade do array =", cap(array))

	fmt.Println("Tamanho do slice =", len(slice))
	fmt.Println("Capacidade do slice =", cap(slice))

	fmt.Println("===========================")

	slice = append(slice, 600)
	fmt.Println("Tamanho do slice =", len(slice))
	fmt.Println("Capacidade do slice =", cap(slice))
}
