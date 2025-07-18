package main

import (
	// Mudando o nome do pacote para System
	System "fmt"
)

func main() {
	System.Printf("Type: %T - Value: %v - Go value - %#v\n", true, true, true)

	System.Printf("Type: %T - Value: %v - Go value - %#v\n", 1, 1, 1)

	System.Printf("Type: %T - Value: %v - Go value - %#v\n", "Arboleda", "Arboleda", "Arboleda")

	System.Printf("Type: %T - Value: %v - Go value - %#v\n", 5.33454, 5.33454, 5.33454)

	System.Printf("Eu quero printar a %% de  %v%%", 34)

	System.Printf("\n\n")

	nome, _, _:= nomeSobrenome1("Arboleda", "de", "Souza")

	System.Printf("Type: %T - Value: %v - Go value - %#v\n", nome, nome, nome)
}

func nomeSobrenome1(nome, meio, sobrenome string) (string, string, string) {
	return nome + "1", meio + "1", sobrenome + "1"
}