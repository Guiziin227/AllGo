package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	dados map[string]any
	mu    sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		dados: make(map[string]any),
	}
}

func (c *Cache) Set(chave string, valor any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.dados[chave] = valor
	fmt.Printf("Valor %x adicionado com a chave %s!\n", valor, chave)
}

func (c *Cache) Get(chave string) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	valor, exists := c.dados[chave]
	return valor, exists

}

func main() {

	var (
		cache = NewCache()
		wg    sync.WaitGroup
	)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cache.Set(fmt.Sprintf("chave%d", i), i)
		}()
	}

	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			time.Sleep(time.Millisecond * 100)
			valor, ok := cache.Get(fmt.Sprintf("chave%d", i))
			if ok {
				fmt.Printf("Valor da chave%d: %v\n", i, valor)
			} else {
				fmt.Printf("Chave%d não encontrada\n", i)
			}
		})
	}

	wg.Wait()
	fmt.Println("Todos os valores foram adicionados e recuperados com sucesso!")
}
