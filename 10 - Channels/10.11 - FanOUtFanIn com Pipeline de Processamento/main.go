package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. Produtor / Gerador:
// Envia todos os números para o canal e fecha o canal ao final.
func gerador(jobs chan<- int, total int) {
	for i := 1; i <= total; i++ {
		jobs <- i
	}
	close(jobs) // Avisa a todos os workers que a lista de tarefas acabou
}

// 2. Worker (Fan-Out):
// Vários workers executam esta função concorrentemente.
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range jobs {
		time.Sleep(20 * time.Millisecond) // Simula custo de processamento

		// Exemplo de filtro: seleciona apenas números pares
		if num%2 == 0 {
			quadrado := num * num
			results <- quadrado
		}
	}
}

func main() {
	totalJobs := 100
	numWorkers := 3

	jobs := make(chan int)
	results := make(chan int)

	// Dispara o gerador em background
	go gerador(jobs, totalJobs)

	// Dispara a pool de workers concorrentes
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// 3. Goroutine de encerramento:
	// Espera todos os workers finalizarem para fechar o canal de resultados com segurança.
	go func() {
		wg.Wait()
		close(results)
	}()

	// 4. Agregador / Redutor (Fan-In no main):
	// A main consome do canal até ele ser fechado e calcula a soma total.
	somaTotal := 0
	itensProcessados := 0

	for res := range results {
		somaTotal += res
		itensProcessados++
	}

	fmt.Printf("Itens filtrados e processados: %d\n", itensProcessados)
	fmt.Printf("Soma total dos quadrados dos pares: %d\n", somaTotal)
}
