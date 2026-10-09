package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	dados sync.Map
}

func NewCache() *Cache {
	return &Cache{
		dados: sync.Map{},
	}
}

func (c *Cache) Set(chave string, valor any) {
	c.dados.Store(chave, valor)
	fmt.Printf("Valor %v adicionado com a chave %s!\n", valor, chave)
}

func (c *Cache) Get(chave string) (any, bool) {
	valor, exists := c.dados.Load(chave)
	return valor, exists

}

func main() {

	var c = NewCache()
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Go(func() {
			c.Set(fmt.Sprintf("chave_%d", i), i)
		})
	}

	for i := 1; i <= 10; i++ {
		wg.Go(func() {
			time.Sleep(1 * time.Second)
			valor, ok := c.Get(fmt.Sprintf("chave_%d", i))
			if ok {
				fmt.Println("O valor: ", valor)
			} else {
				fmt.Println("Valor não encontrado")
			}
		})
	}

	wg.Wait()

}
