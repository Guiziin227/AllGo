package main

import (
	"fmt"
	"strings"
	"time"
)

// Transforma tudo em Maiúsculas
func consumidorMaiusculo(ch <-chan string) {
	for valor := range ch {
		fmt.Printf("[Maiúsculas]: %s\n", strings.ToUpper(valor))
	}
}

// Transforma tudo em Minúsculas
func consumidorMinusculo(ch <-chan string) {
	for valor := range ch {
		fmt.Printf("[Minúsculas]: %s\n", strings.ToLower(valor))
	}
}

// Mantém apenas as vogais
func consumidorApenasVogais(ch <-chan string) {
	for valor := range ch {
		filtrado := ""
		vogais := "aeiouAEIOU"
		for _, caractere := range valor {
			if strings.ContainsRune(vogais, caractere) {
				filtrado += string(caractere)
			}
		}
		fmt.Printf("[Só Vogais]: %s\n", filtrado)
	}
}

func broadcast(channels []chan string, message string) {
	for _, channel := range channels {
		channel <- message
	}

	for _, channel := range channels {
		close(channel)
	}
}

func main() {
	// Criando os 3 canais de string
	var channels = []chan string{
		make(chan string),
		make(chan string),
		make(chan string),
	}

	go consumidorMaiusculo(channels[0])
	go consumidorMinusculo(channels[1])
	go consumidorApenasVogais(channels[2])

	// Enviando a mensagem para todos os canais
	broadcast(channels, "golanguinho")

	// Espera as goroutines terminarem de processar e printar
	time.Sleep(100 * time.Millisecond)
}
