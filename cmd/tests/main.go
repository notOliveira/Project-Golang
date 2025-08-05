package main

import (
	// Mudando o nome do pacote para System
	"fmt"
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

	type Pessoa struct {
		Nome      string
		Sobrenome string
		Idade     int
	}

	var gustavo = Pessoa{
		Nome:	"Gustavo",
		Sobrenome: "Arboleda",
		Idade:     30,
	}

	var pessoasMap =  make(map[int]Pessoa)

	fmt.Printf("Nome: %v - Sobrenome: %v - Idade: %v\n\n", gustavo.Nome, gustavo.Sobrenome, gustavo.Idade)
	
	for i := 0; i < 20; i++ {
		pessoasMap[i] = Pessoa{
			Nome:      fmt.Sprintf("Aluno %d", i),
			Sobrenome: fmt.Sprintf("Fulano %d", i),
			Idade:     i + 20,
		}
	}

	for i, pessoa := range pessoasMap {
		System.Printf("Index: %v - Nome: %v - Sobrenome: %v - Idade: %v\n", i, pessoa.Nome, pessoa.Sobrenome, pessoa.Idade)
	}

	if idade := 18; idade >= 18 {
        fmt.Println("Maior de idade")
    } else {
        fmt.Println("Menor de idade")
    }

}
func nomeSobrenome1(nome, meio, sobrenome string) (string, string, string) {
	return nome + "1", meio + "1", sobrenome + "1"
}