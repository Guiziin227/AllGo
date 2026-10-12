package main

import (
	"fmt"
	"sync"
)

func produtor(entrada chan<- int, numeros []int) {
	for _, num := range numeros {
		entrada <- num
	}
	close(entrada)
}

func fanOut(id int, entrada <-chan int, saida chan<- int) {
	for num := range entrada {
		fmt.Printf("Channel %d processando o numero %d\n", id, num)
		saida <- num * num
	}
	close(saida)
}

func fanIn(saidasReadOnly []<-chan int) <-chan int {
	saidaFinal := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(saidasReadOnly))

	for _, saida := range saidasReadOnly {
		go func() {
			defer wg.Done()

			for num := range saida {
				saidaFinal <- num
			}
		}()
	}

	go func() {
		wg.Wait()
		close(saidaFinal)
	}()

	return saidaFinal
}

func main() {

	var (
		numeros = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // workload

		entrada = make(chan int)
		saidas  = []chan int{
			make(chan int),
			make(chan int),
			make(chan int),
			make(chan int),
		}
	)

	go produtor(entrada, numeros)

	for i, saida := range saidas {
		go fanOut(i+1, entrada, saida)
	}

	saidasReadOnly := make([]<-chan int, len(saidas)) // na minha lista só aceito channels de receive
	for i, saida := range saidas {
		saidasReadOnly[i] = saida
	}

	saidaFinal := fanIn(saidasReadOnly)
	for num := range saidaFinal {
		fmt.Printf("Resultado final: %d\n", num)
	}
}
