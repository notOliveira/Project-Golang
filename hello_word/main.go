package main

import (
	"fmt"
)

func exit() {
	fmt.Print("\nPressione ENTER para sair...")
	fmt.Scanln()
}

func main() {
	fmt.Println("Hello world!")

	exit()
}
