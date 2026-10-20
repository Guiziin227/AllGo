package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

func semPool(iteracoes int) time.Duration {

	start := time.Now()
	for i := 0; i < iteracoes; i++ {
		buffer := bytes.Buffer{}
		buffer.WriteString("Olá. mundo")
		_ = buffer.String()
	}
	return time.Since(start)
}

func comPool(iteracoes int) time.Duration {
	start := time.Now()

	pool := sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}

	for i := 0; i < iteracoes; i++ {
		buffer := pool.Get().(*bytes.Buffer)

		buffer.WriteString("Olá, mundo")
		_ = buffer.String()

		buffer.Reset()
		pool.Put(buffer)
	}

	return time.Since(start)
}

func porcentagem(t1, t2 time.Duration) float64 {
	return (float64(t1-t2) / float64(t1)) * 100
}

func main() {

	iteracoes := 1_000_000_000

	fmt.Printf("Executando comp com %d iteracoes\n", iteracoes)

	t1 := semPool(iteracoes)
	fmt.Printf("[Sem pool] Tempo total: %s\n", t1)
	t2 := comPool(iteracoes)
	fmt.Printf("[Com pool] Tempo total: %s\n", t2)

	fmt.Printf("Porcentagem de melhoria: %.2f%%\n", porcentagem(t1, t2))
}

// Resultado
/*
Executando comp com 1000000000 iteracoes
[Sem pool] Tempo total: 1m5.353834252s
[Com pool] Tempo total: 19.569810165s
Porcentagem de melhoria: 70.06%
*/
