package main

import (
	"bytes"
	"fmt"
	"sync"
)

var pool = sync.Pool{
	New: func() any {
		fmt.Printf("Criando novo buffer...\n")
		return new(bytes.Buffer)
	},
}

// com o sync pool a gente faz menos alocações de memória e o gc seja um pouco mais fácil
func main() {

	buffer := pool.Get().(*bytes.Buffer)
	buffer.WriteString("Escrevendo...")
	fmt.Println(buffer.String())

	buffer.Reset()
	pool.Put(buffer)

	buffer2 := pool.Get().(*bytes.Buffer)
	buffer2.WriteString("Escrevendo 2...")
	fmt.Println(buffer2.String())

}
