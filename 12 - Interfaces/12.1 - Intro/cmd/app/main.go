package main

import (
	"github.com/guiziin227/allgo/interfaces/intro/pkg/conversor"
	"github.com/guiziin227/allgo/interfaces/intro/pkg/conversor/json"
	"github.com/guiziin227/allgo/interfaces/intro/pkg/conversor/xml"
)

func main() {

	gerenciador := conversor.NewGerenciador()
	gerenciador.RegistrarConversor(json.ConversorJSON{})
	gerenciador.RegistrarConversor(xml.ConversorXML{})

	testeJSON := `{
						"pessoa":{
						"nome": "Alice", "idade": 30}}`

	testeXML := `<pessoa><nome>Alice</nome><idade>30</idade></pessoa>`
	gerenciador.Processar("XML", testeXML)
	gerenciador.Processar("JSON", testeJSON)
}
