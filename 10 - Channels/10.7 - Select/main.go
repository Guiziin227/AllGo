package main

import "time"

func enviarMensagem(ch chan<- string, msg string, intervalo time.Duration) {
	for {
		time.Sleep(intervalo)
		ch <- msg
	}
}

func main() {

	var (
		ch1 = make(chan string)
		ch2 = make(chan string)
	)

	go enviarMensagem(ch1, "Mensagem 1", 1*time.Second)
	go enviarMensagem(ch2, "Mensagem 2", 2*time.Second)

	for range 10 {
		select {
		case msg := <-ch1:
			println(msg)
		case msg := <-ch2:
			println(msg)
		case <-time.After(3 * time.Second):
			println("timeout")
			return
		}
	}
}
