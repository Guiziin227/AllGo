package main

import "fmt"

func main() {
	var a int = 10
	var p *int = &a

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(p)
	fmt.Println(*p)

	var zeroValue *int
	var zeroValue2 *string

	fmt.Println(zeroValue)
	fmt.Println(zeroValue2)

	fmt.Println("========criando pessoa========")

	pessoa := NewPessoa("Gui", 20)
	fmt.Println(pessoa.Nome, pessoa.Idade, pessoa.Telefone())
	pessoa.AtualizarIdade(25)
	fmt.Println(pessoa.Nome, pessoa.Idade, pessoa.Telefone())

	pessoa.AtualizarTelefone("55991654341")
	fmt.Println(pessoa.Nome, pessoa.Idade, pessoa.Telefone())
}

type Pessoa struct {
	Nome     string
	Idade    int
	telefone *string
}

func NewPessoa(nome string, idade int) Pessoa {
	return Pessoa{
		Nome:  nome,
		Idade: idade,
	}
}

// Telefone faz com que evite o panic por não existir nada no ponteiro
// Escondemos os campos que são ponteiros e tratamos dentro da função
func (p *Pessoa) Telefone() string {
	if p.telefone == nil {
		return ""
	}
	return *p.telefone
}

// AtualizarIdade atualiza a idade da pessoa
func (p *Pessoa) AtualizarIdade(idade int) {
	p.Idade = idade
}

// AtualizarTelefone atualiza o telefone da pessoa
func (p *Pessoa) AtualizarTelefone(telefone string) {
	p.telefone = &telefone
}
