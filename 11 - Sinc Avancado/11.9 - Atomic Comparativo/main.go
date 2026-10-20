package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func testeMutex(goroutines, incrementos int) (int64, time.Duration) {
	var (
		mutex    sync.Mutex
		contador int64
		start    = time.Now()
		wg       sync.WaitGroup
	)

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < incrementos; j++ {
				mutex.Lock()
				contador++
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()
	return contador, time.Since(start)
}

func testeAtomic(goroutines, incrementos int) (int64, time.Duration) {
	var (
		contador int64
		start    = time.Now()
		wg       sync.WaitGroup
	)

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < incrementos; j++ {
				atomic.AddInt64(&contador, 1)
			}
		}()
	}
	wg.Wait()
	return contador, time.Since(start)
}

func porcentagem(t1, t2 time.Duration) float64 {
	return (float64(t1-t2) / float64(t1)) * 100
}

func main() {

	iteracoes := 1_000_000
	goroutines := 1_000

	c1, t1 := testeMutex(goroutines, iteracoes)
	fmt.Printf("[Mutex] Contador final: %d, tempo: %v\n", c1, t1)

	c2, t2 := testeAtomic(goroutines, iteracoes)
	fmt.Printf("[Atomic] Contador final: %d, tempo: %v\n", c2, t2)

	fmt.Printf("[Resultado] %.2f%% mais rápido\n", porcentagem(t1, t2))
}

// Resultado:
/*
	[Mutex] Contador final: 1000000000, tempo: 36.510904908s
	[Atomic] Contador final: 1000000000, tempo: 7.601157716s
	[Resultado] 79.18% mais rápido
*/
