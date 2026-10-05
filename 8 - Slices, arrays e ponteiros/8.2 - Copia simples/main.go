package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1 // Isso cria uma cópia da referência, não do conteúdo

	fmt.Println(slice1, slice2)

	slice2 = append(slice2, 4)
	slice2[0] = 10

	fmt.Println(slice1, slice2)

	fmt.Println("==================")

	var slice3 = []int{1, 2, 3, 4, 5}
	var slice4 = make([]int, len(slice3))

	copy(slice4, slice3)

	slice3[0] = 100

	slice4 = append(slice4, 2)
	fmt.Println(slice3, slice4)

}
