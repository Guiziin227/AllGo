package main

import "fmt"

func main() {

	// Estrutura chave e valor

	var meuMap = map[string]any{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
	}

	for chave, valor := range meuMap {
		fmt.Println(chave, valor)
	}

	fmt.Println("====================")

	var meuMap2 = map[string]any{}

	meuMap2["a"] = "oie"

	var valor, ok = meuMap2["b"].(string)

	if ok {
		fmt.Println(valor)
	} else {
		fmt.Println("key not found")
	}

}
