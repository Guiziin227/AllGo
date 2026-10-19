package main

import (
	"fmt"
	"sync"
	"time"
)

type Trava struct {
	liberada bool
	mu       sync.Mutex
	cond     *sync.Cond
}

func NovaTrava() *Trava {
	t := &Trava{}

	t.cond = sync.NewCond(&t.mu)
	return t

}

var start = time.Now()

func logf(msg string) {
	fmt.Printf("[%4dms] %s\n", time.Since(start).Milliseconds(), msg)
}

func worker(id int, t *Trava, wg *sync.WaitGroup) {
	defer wg.Done()

	t.mu.Lock()

	defer t.mu.Unlock()

	for !t.liberada {
		logf(fmt.Sprintf("Worker %d is waiting", id))
		t.cond.Wait()
	}

	logf(fmt.Sprintf("Worker %d is done", id))
	time.Sleep(1 * time.Second)
}

func main() {

	{
		var (
			qtdWorkers = 4
			trava      = NovaTrava()
			wg         sync.WaitGroup
		)

		wg.Add(qtdWorkers)

		for i := range qtdWorkers {
			go worker(i, trava, &wg)
		}

		time.Sleep(1 * time.Second)

		trava.mu.Lock()
		trava.liberada = true
		logf("Trava liberada!")
		trava.cond.Broadcast()
		trava.mu.Unlock()

		wg.Wait()
	}

	{
		var (
			qtdWorkers = 4
			trava      = NovaTrava()
			wg         sync.WaitGroup
		)

		wg.Add(qtdWorkers)

		for i := range qtdWorkers {
			go worker(i, trava, &wg)
		}

		time.Sleep(1 * time.Second)

		trava.mu.Lock()
		trava.liberada = true
		logf("Trava liberada!")
		trava.cond.Signal()
		trava.mu.Unlock()

		wg.Wait()
	}

}
