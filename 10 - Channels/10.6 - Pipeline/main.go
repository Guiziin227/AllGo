package main

import "fmt"

func etapa1(ch1 chan<- int, numeros []int) {
	for _, n := range numeros {
		ch1 <- n
	}
	close(ch1)
}

func etapa2(ch1 <-chan int, ch2 chan<- int) {
	for c1 := range ch1 {
		ch2 <- c1 * c1
	}
	close(ch2)
}

func etapa3(ch2 <-chan int) {
	for c2 := range ch2 {
		fmt.Printf("Valor recebido na etp3: %d\n ", c2)
	}
}

func main() {

	var (
		numeros = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		ch1     = make(chan int)
		ch2     = make(chan int)
	)

	go etapa1(ch1, numeros)
	go etapa2(ch1, ch2)
	etapa3(ch2)
}
