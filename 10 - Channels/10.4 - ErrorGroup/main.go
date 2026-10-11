package main

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d iniciado\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d finalizado\n", id)

}

func workerErrGroup(id int) error {
	fmt.Printf("worker %d iniciado\n", id)

	time.Sleep(time.Second)

	if id == 1 {
		return errors.New("erro inesperado no worker")
	}
	fmt.Printf("worker %d finalizado\n", id)
	return nil
}

func main() {

	fmt.Println("==== Wait Group ====")

	var wg sync.WaitGroup

	for i := range 3 {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

	fmt.Println("Foi finalizado")

	fmt.Println("==== Error Group ====")

	var eg errgroup.Group

	eg.SetLimit(3)

	for i := range 5 {
		eg.Go(func() error {
			return workerErrGroup(i)
		})
	}

	// Pontos importantes:
	// 1 - Se duas ou mais goroutiones derem erro, o erro mostrado só será da primeira
	// 2 - O ErrorGroup não interrompe as outras goroutines quando uma delas falha - isso pode ser contornado com 'context'
	err := eg.Wait()
	if err != nil {
		fmt.Println("Erro encontrado:", err.Error())
	} else {
		fmt.Println("Nenhum erro encontrado")
	}
}
