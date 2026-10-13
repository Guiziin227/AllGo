package main

import (
	"fmt"
	"sync"
	"time"
)

type Buffer struct {
	dados []int
	cond  *sync.Cond
	mu    sync.Mutex
}

func main() {

	buffer := Buffer{
		dados: make([]int, 0),
	}

	buffer.cond = sync.NewCond(&buffer.mu)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := range 5 {
			buffer.mu.Lock()
			fmt.Println("Produzindo item ", i)
			buffer.dados = append(buffer.dados, i)
			buffer.cond.Signal()
			buffer.mu.Unlock()
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {

		defer wg.Done()

		for range 5 {
			buffer.mu.Lock()
			for len(buffer.dados) == 0 {
				buffer.cond.Wait()
			}
			item := buffer.dados[0]
			buffer.dados = buffer.dados[1:]
			fmt.Println("Consumindo item ", item)
			buffer.mu.Unlock()
		}
	}()
	wg.Wait()
}
