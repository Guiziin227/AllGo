package main

import (
	"fmt"
	"time"
)

func enviar(ch chan int) {
	fmt.Println("Enviando valor para o channel")
	ch <- 100 // ao enviar ele trava esperando que alguém receba
	fmt.Println("Valor enviado com sucesso")
}

func receber(ch chan int) {
	time.Sleep(time.Second)
	valor := <-ch
	fmt.Println("Recebido: ", valor)
}

// Channels são seguros de serem compartilhados entre goroutines
func main() {

	// Comunicação entre go routines. Use channels.
	ch := make(chan int) // Channel sem buffer

	for i := 0; i < 3; i++ {
		go enviar(ch)
	}

	time.Sleep(time.Second)

	for i := 0; i < 3; i++ {
		go receber(ch)
	}

	time.Sleep(5 * time.Second)

	//Cuidar com deadlocks
	<-ch
}
