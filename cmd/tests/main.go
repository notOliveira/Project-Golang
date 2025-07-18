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

	s := make([]int, 5)

	s[0] = 3214234
	s[1] = 3
	s[2] = 233
	
	System.Println(s)

	var s1 []int
	s1 = append(s1, 1)

	System.Println(s1)

	var ar [3]int = [3]int{2}
	ar[2] = 4

	System.Println(ar)

}
func nomeSobrenome1(nome, meio, sobrenome string) (string, string, string) {
	return nome + "1", meio + "1", sobrenome + "1"
}