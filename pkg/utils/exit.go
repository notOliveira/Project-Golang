package utils

import (
	"bufio"
	"fmt"
	"os"
)

// Função exit que aguarda o usuário pressionar ENTER
func Exit() {
	fmt.Print("Pressione ENTER para sair...")
	bufio.NewReader(os.Stdin).ReadString('\n') // Aguarda o ENTER
}
