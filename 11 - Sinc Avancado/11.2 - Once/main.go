package main

import (
	"fmt"
	"sync"
	"time"
)

type BancoDeDados struct {
	conectado bool
}

var (
	bancoDeDados *BancoDeDados
	once         sync.Once
)

func ObterBancoDeDados() *BancoDeDados {
	once.Do(func() {
		fmt.Println("Conectando no banco de dados")
		time.Sleep(time.Second * 2)
		bancoDeDados = &BancoDeDados{conectado: true}
		fmt.Println("Conectado!")
	})
	return bancoDeDados
}

func main() {

	var (
		wg sync.WaitGroup
	)

	wg.Add(5)

	for i := range 5 {
		go func(n int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d tentando acessar o banco.\n", n+1)
			banco := ObterBancoDeDados()
			fmt.Printf("Goroutine %d acessou o banco. Conectado=%v\n", n+1, banco.conectado)
		}(i)
	}

	wg.Wait()

	fmt.Println("Finalizando...")
}
