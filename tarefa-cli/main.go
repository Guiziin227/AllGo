package main

import (
	"bufio"
	"fmt"

	"os"
	"strconv"
	"strings"

	"github.com/guiziin227/allgo/tarefacli/task"
)

func main() {

	manager := task.NewManager()
	buf := bufio.NewScanner(os.Stdin)

	err := manager.Load("tasks.json")
	if err != nil {
		return
	}

	for {
		fmt.Println("Escolha uma opção:")

		buf.Scan()
		if err := buf.Err(); err != nil {
			fmt.Println("Erro ao ler entrada:", err)
			return
		}
		op := buf.Text()
		args := strings.Fields(op)

		if len(args) < 1 {
			fmt.Println("Opção inválida")
			continue
		}

		switch args[0] {
		case "add":
			fmt.Println("Adicionando tarefa...")
			_, err := manager.Add(strings.Join(args[1:], " "))
			if err != nil {
				fmt.Println("Erro ao adicionar tarefa:", err)
			}
		case "list":
			fmt.Println("Listando tarefas...")
			tasks := manager.List()
			for _, t := range tasks {
				fmt.Println(t)
			}
		case "complete":
			if len(args) < 2 {
				fmt.Println("Opção inválida")
				continue
			}
			id, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Erro ao converter ID:", err)
				continue
			}
			err = manager.Complete(id)
			if err != nil {
				fmt.Println("Erro ao completar tarefa:", err)
			}
			fmt.Println("Completando tarefa...")
		case "exit":
			fmt.Println("Finalizando programa...")
			err := manager.Save("tasks.json")
			if err != nil {
				fmt.Println("Erro ao salvar tarefas:", err)
			}
			return
		default:
			fmt.Println("Opção inválida")
			fmt.Println("Temos 'add', 'list', 'complete', 'exit'")
		}

	}

}
