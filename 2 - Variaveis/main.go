package main

import "fmt"

func main() {

	var long string
	long = "Variavel Longa"

	short := "Variavel Curta"

	fmt.Println(long, short)

	var (
		variavel  string = "i"
		variavel2 string = "o"
	)

	fmt.Println(variavel, variavel2)

	variavel3, variavel4 := "a", "e"

	fmt.Println(variavel3, variavel4)
}
