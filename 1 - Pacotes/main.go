package main

import (
	"fmt"
	"modulo/auxiliar"
	"modulo/externo"
)

func main() {

	fmt.Println("Main")

	auxiliar.Mostrando()

	err := externo.ValidandoEmail("exemplo@email.com")
	if err {
		fmt.Println("Email válido")
	} else {
		fmt.Println("Email inválido")
	}

}
