package main

import (
	"context"
	"fmt"
)

type (
	myString1 string
	myString2 string
)

// O context em GO é uma ferramenta para passar informações compartilhadas
// e controlar o ciclo de vida de operações em um programa

func main() {

	ctx := context.Background()
	ctx = context.WithValue(ctx, myString1("key"), "value1") // essa funcao retorna um context novo...
	ctx = context.WithValue(ctx, myString2("key"), "value2")

	leitura(ctx)
}

func leitura(ctx context.Context) {
	valoArmazenado := ctx.Value(myString2("key"))
	if valoArmazenado != nil {
		fmt.Println("Valor encontrado:", valoArmazenado)
	} else {
		fmt.Println("Valor não encontrado")
	}

}
