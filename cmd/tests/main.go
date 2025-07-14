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
}
