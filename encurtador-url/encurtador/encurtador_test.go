package encurtador

import (
	"sync"
	"testing"
)

func TestConcurrentResolveAndStats(t *testing.T) {
	s := NewShortener()

	created, err := s.Shorten("https://www.google.com")
	if err != nil {
		t.Fatalf("erro ao criar link: %v", err)
	}

	const totalGoroutines = 100
	var wg sync.WaitGroup

	// Dispara 100 acessos simultâneos (escrita: Clicks++)
	// e 100 leituras de estatísticas simultâneas
	for i := 0; i < totalGoroutines; i++ {
		wg.Add(2)

		// Goroutine 1: Simula alguém acessando o link curto
		go func() {
			defer wg.Done()
			_, err := s.Resolve(created.Code)
			if err != nil {
				t.Errorf("erro inesperado no Resolve: %v", err)
			}
		}()

		// Goroutine 2: Simula alguém consultando o painel de estatísticas
		go func() {
			defer wg.Done()
			_, err := s.Stats(created.Code)
			if err != nil {
				t.Errorf("erro inesperado no Stats: %v", err)
			}
		}()
	}

	// Espera todas as 200 goroutines terminarem
	wg.Wait()

	// Verifica se nenhum clique se perdeu no meio do caminho!
	stats, err := s.Stats(created.Code)
	if err != nil {
		t.Fatalf("erro ao buscar stats finais: %v", err)
	}

	if stats.Clicks != totalGoroutines {
		t.Errorf("esperava exatos %d cliques, mas contou %d (houve perda por concorrência!)", totalGoroutines, stats.Clicks)
	}
}
