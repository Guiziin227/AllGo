package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var (
		flag int64
		wg   sync.WaitGroup
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Escrevendo...")
		time.Sleep(2 * time.Second)
		atomic.StoreInt64(&flag, 1)
	}()

	go func() {
		defer wg.Done()
		for atomic.LoadInt64(&flag) == 0 {
			fmt.Println("Flag com o valor 0...")
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("Flag com o valor 1...")
	}()

	wg.Wait()
}
