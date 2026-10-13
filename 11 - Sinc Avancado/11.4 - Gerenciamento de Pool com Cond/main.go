package main

import (
	"fmt"
	"sync"
	"time"
)

type PoolDeConexoes struct {
	conexoes []int
	mu       sync.Mutex
	cond     *sync.Cond
}

func NewPoolDeConexoes(tam int) *PoolDeConexoes {
	pool := &PoolDeConexoes{
		conexoes: make([]int, tam),
	}

	for i := range tam {
		pool.conexoes[i] = i + 1
	}

	pool.cond = sync.NewCond(&pool.mu)
	return pool
}

func (p *PoolDeConexoes) ObterConexo() int {
	p.mu.Lock()
	defer p.mu.Unlock()

	for len(p.conexoes) == 0 {
		fmt.Println("Nenhuma conexao disponivel. Aguardando...")
		p.cond.Wait()
	}

	conexao := p.conexoes[0]
	p.conexoes = p.conexoes[1:]
	fmt.Printf("Conexão %d adquirida\n", conexao)
	return conexao
}

func (p *PoolDeConexoes) LiberarConexao(conexao int) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.conexoes = append(p.conexoes, conexao)
	fmt.Printf("Conexao %d liberada.\n", conexao)
	p.cond.Signal()
}

func main() {

	poolDeConexoes := NewPoolDeConexoes(2)

	var wg sync.WaitGroup
	wg.Add(5)

	for i := range 5 {
		go func() {
			defer wg.Done()

			conexao := poolDeConexoes.ObterConexo()
			fmt.Printf("Goroutine %d adquiriu a conexao %d\n", i, conexao)
			time.Sleep(2 * time.Second)
			poolDeConexoes.LiberarConexao(conexao)
		}()
	}

	wg.Wait()
}
