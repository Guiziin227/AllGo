package main

import (
	"fmt"
	"sync"
)

func main() {

	var (
		m  sync.Map
		wg sync.WaitGroup
	)

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			m.Store(i, i)
		}()
	}

	wg.Wait()

	m.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})

	valor, ok := m.Load(20)
	if ok {
		fmt.Println("Valor da chave 20:", valor)
	} else {
		fmt.Println("Chave 20 não encontrada")
	}
}
