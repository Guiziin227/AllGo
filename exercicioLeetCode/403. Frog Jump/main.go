package main

import "fmt"

func main() {
	stone := []int{0, 1, 3, 5, 6, 8, 17}

	fmt.Println(canCross(stone))
}

func canCross(stones []int) bool {
	if len(stones) == 0 {
		return false
	}

	// Guarda quais tamanhos de salto podem chegar em cada pedra.
	saltos := make([]map[int]bool, len(stones))

	for i := range saltos {
		saltos[i] = make(map[int]bool)
	}

	// Na pedra inicial, ainda não houve salto.
	saltos[0][0] = true

	for i := 0; i < len(stones); i++ {
		for tamPulo := range saltos[i] {

			// O próximo salto pode ser k-1, k ou k+1.
			for proximoPulo := tamPulo - 1; proximoPulo <= tamPulo+1; proximoPulo++ {
				if proximoPulo <= 0 {
					continue
				}

				proximaPosicao := stones[i] + proximoPulo

				for j := i + 1; j < len(stones); j++ {
					if stones[j] == proximaPosicao {
						saltos[j][proximoPulo] = true
						break
					}

					if stones[j] > proximaPosicao {
						break
					}
				}
			}
		}
	}

	return len(saltos[len(stones)-1]) > 0
}
