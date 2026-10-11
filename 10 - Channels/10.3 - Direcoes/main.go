package main

import (
	"fmt"
	"time"
)

func enviar(ch chan<- int) {
	fmt.Println("Enviando valor para o channel")
	ch <- 100
	fmt.Println("Valor enviado com sucesso")
}

func receber(ch <-chan int) {
	time.Sleep(time.Second)
	valor := <-ch
	fmt.Println("Recebido: ", valor)
}

func main() {

	ch := make(chan int, 3)

	for range 5 {
		go enviar(ch)
	}

	time.Sleep(time.Second)

	for range 5 {
		go receber(ch)
	}

	time.Sleep(5 * time.Second)
}
