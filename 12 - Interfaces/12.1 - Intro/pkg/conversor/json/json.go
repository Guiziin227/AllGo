package json

import "encoding/json"

type ConversorJSON struct{}

func (j ConversorJSON) Converter(dados string) (map[string]any, error) {
	var resultado map[string]any
	err := json.Unmarshal([]byte(dados), &resultado)
	return resultado, err
}

func (j ConversorJSON) Formato() string {
	return "JSON"
}
