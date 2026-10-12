package main

import (
	"fmt"
	"sync"
	"time"
)

//
//func main() {
//	ch := make(chan int)
//
//	numero := 7
//
//	go func(n int) {
//		dobro := n * 2
//		ch <- dobro
//	}(numero)
//	resultado := <-ch
//	fmt.Printf("O dobro de %d é %d\n", numero, resultado)
//}

//func main() {
//
//	var wg sync.WaitGroup
//
//	for i := range 3 {
//		wg.Add(1)
//		go func(id int) {
//			defer wg.Done()
//			fmt.Println("Iniciando a tarefa: ", id)
//			time.Sleep(1 * time.Second)
//			fmt.Println("Finalizando a tarefa: ", id)
//		}(i)
//	}
//
//	wg.Wait()
//	fmt.Println("Todas as tasks concluidas")
//}

//func main() {
//
//	var wg sync.WaitGroup
//	numeros := []int{2, 4, 6}
//	resultados := make(chan int, 3)
//
//	for _, numero := range numeros {
//		wg.Add(1)
//		go func(n int) {
//			defer wg.Done()
//			time.Sleep(1 * time.Second)
//
//			resultados <- n * n
//		}(numero)
//
//	}
//
//	wg.Wait()
//	close(resultados)
//
//	fmt.Println("Resultados: ", <-resultados, <-resultados, <-resultados)
//
//}

func main() {
	palavras := []string{"golang", "concorrencia", "goroutine", "threads", "canais"}

	var wg sync.WaitGroup

	jobs := make(chan string, len(palavras))
	results := make(chan string, len(palavras))
	for _, palavra := range palavras {
		jobs <- palavra
	}
	close(jobs)

	workers := 2
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			for palavra := range jobs {
				time.Sleep(500 * time.Millisecond)
				results <- fmt.Sprintf("Processando: %s (worker %d)", palavra, w)
			}
		}(i)
	}

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println(result)
	}

}
