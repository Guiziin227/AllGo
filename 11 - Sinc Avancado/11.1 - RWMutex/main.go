package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	dados map[string]any
	mu    sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		dados: make(map[string]any),
	}
}

func (c *Cache) Set(chave string, valor any) {
	logf("Set: Tentativa de escrita (Lock)")
	c.mu.Lock()
	logf("Set: Lock Adquirido")
	time.Sleep(300 * time.Millisecond)
	c.dados[chave] = valor
	c.mu.Unlock()
	logf("Set: Unlock (escrita finalizada)")
}

func (c *Cache) Get(chave string) (any, bool) {
	logf("Get: Tentativa de leitura (RLock)")
	c.mu.RLock()
	logf("Get: RLock Adquirido")
	time.Sleep(100 * time.Millisecond)
	valor, exists := c.dados[chave]
	c.mu.RUnlock()
	logf("Get: RUnlock (leitura finalizada)")
	return valor, exists
}

var start = time.Now()

func logf(msg string) {
	fmt.Printf("[%4dms] - %s\n", time.Since(start).Milliseconds(), msg)
}

func main() {
	var (
		cache = NewCache()
		wg    sync.WaitGroup
	)

	wg.Add(4)
	cache.dados["chave_1"] = "1"
	cache.dados["chave_2"] = "2"

	go func() {
		defer wg.Done()
		logf("Goroutine 1: chama get")
		_, _ = cache.Get("chave_1")
		logf("Goroutine 1: recebeu get")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		logf("Goroutine 2: chama get")
		_, _ = cache.Get("chave_1")
		logf("Goroutine 2: recebeu get")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(20 * time.Millisecond)
		logf("Goroutine 3: chama set")
		cache.Set("chave_1", "3")
		logf("Goroutine 3: recebeu set")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(300 * time.Millisecond)
		logf("Goroutine 4: chama get")
		_, _ = cache.Get("chave_1")
		logf("Goroutine 4: recebeu get")
	}()

	wg.Wait()

	logf("Fim")
}
