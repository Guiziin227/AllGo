package main

import (
	"sync"
)

func main() {

	var (
		m  = make(map[int]int)
		mu sync.Mutex
		wg sync.WaitGroup
	)

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()

			mu.Lock() // Bloqueia o acesso ao mapa para evitar condições de corrida
			m[i] = i
			mu.Unlock()
		}()
	}

	wg.Wait()

	for c, v := range m {
		println(c, v)
	}
}
