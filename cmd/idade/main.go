package main

import (
	"fmt"
	"projectGO/pkg/utils"
)

func exit() {
	fmt.Print("\nPressione ENTER para sair...")
	fmt.Scanln()
}

func main() {
	var idade int

	fmt.Print("Digite sua idade: ")
	_, err := fmt.Scan(&idade)
	if err != nil {
		fmt.Println("Erro ao ler a idade. Certifique-se de digitar um número.")
		exit()
		return
	}

	fmt.Printf("\nVocê tem %d anos!\n", idade)

	utils.Exit()
}
