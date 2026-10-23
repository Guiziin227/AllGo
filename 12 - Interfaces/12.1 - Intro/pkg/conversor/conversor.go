package conversor

import "fmt"

type Conversor interface {
	Converter(string) (map[string]any, error)
	Formato() string
}

type Gerenciador struct {
	conversores map[string]Conversor
}

func NewGerenciador() *Gerenciador {
	return &Gerenciador{conversores: make(map[string]Conversor)}
}
func (g *Gerenciador) Processar(formato, dados string) {
	conversor, ok := g.conversores[formato]
	if !ok {
		fmt.Printf("Nenhum conversor registrado para o formato: %s\n", formato)
		return
	}

	resultado, err := conversor.Converter(dados)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Resultado: %v\n", resultado)
}

func (g *Gerenciador) RegistrarConversor(c Conversor) {
	g.conversores[c.Formato()] = c
}
