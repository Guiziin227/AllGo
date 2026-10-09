package main

import "time"

func main() {

	// Só acontece se tiver goroutines

	// duas ou mais goroutines acessando o mesmo dado ao mesmo tempo
	// e ao menos uma delas está modificando o dado

	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}

	m2 := make(map[int]int)

	go func() {
		for i := 0; i < 1000; i++ {
			m2[i] = i
		}
	}()

	go func() {
		for i := 1000; i < 2000; i++ {
			m2[i] = i
		}
	}()

	time.Sleep(5 * time.Second)
}
