package main

import (
	"fmt"
	"time"
)

func agendar(intervalo time.Duration, f func(), channelDeParada <-chan struct{}, channelDeFinalizacao chan<- struct{}) {
	ticker := time.NewTicker(intervalo)
	defer close(channelDeFinalizacao)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			f()
		case <-channelDeParada:
			fmt.Println("sinal de parada recebido!")
			return
		}
	}

}

func main() {

	channelDeParada := make(chan struct{})      //stop
	channelDeFinalizacao := make(chan struct{}) //done

	go agendar(3*time.Second, func() {
		fmt.Println("Executando tarefa...")
	}, channelDeParada, channelDeFinalizacao)

	time.Sleep(10 * time.Second)

	close(channelDeParada)
	<-channelDeFinalizacao
}
