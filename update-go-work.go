package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Encontra todas as pastas com go.mod
	var dirs []string
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == "go.mod" {
			dirs = append(dirs, filepath.Dir(path))
		}
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao caminhar pelo projeto: %v\n", err)
		os.Exit(1)
	}

	if len(dirs) == 0 {
		fmt.Println("Nenhum go.mod encontrado.")
		return
	}

	// Decide entre init ou use
	var cmd *exec.Cmd
	if _, err := os.Stat("go.work"); os.IsNotExist(err) {
		// Se não existir, inicializa o workspace com todos os módulos
		args := append([]string{"work", "init"}, dirs...)
		cmd = exec.Command("go", args...)
		fmt.Printf("Criando novo go.work: go %v\n", args)
	} else {
		// Se já existir, adiciona módulos ao workspace
		args := append([]string{"work", "use"}, dirs...)
		cmd = exec.Command("go", args...)
		fmt.Printf("Atualizando go.work: go %v\n", args)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao executar go work: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("go.work atualizado com sucesso!")
}
