# Anotações sobre Golang. 

<em>Não espere coisas boas daqui, pois estou aprendendo!</em>

## Tipos principais de dados

<table>
    <tr>
        <td><b>bool</b></td>
        <td><b>string</b></td>
        <td><b>int</b></td>
        <td><b>float (float64/float32)</b></td>
    </tr>
        <tr>
        <td>true ou False</td>
        <td>Texto - Sequência de bytes</td>
        <td>Números</td>
        <td>Números decimais</td>
    </tr>
</table>

## Variáveis & Constantes

<details>
    <summary>
    <h3>Declaração rápida de variável (inferência de tipo)</h3>
    </summary>

    idade := 30
</details>

<i>OBS: Só pode ser usado dentro de funções. Só funciona para variáveis</i>

<details>
    <summary>
        <h3>Declaração explícita de variável/constante</h3>
    </summary>
    
    var idade1 int = 30

    const idade2 int = 25

    // Ou

    var nome1 string
    nome = "João"

    // Ou

    var dia = "Terça-Feira"

    const data = "26 de fevereiro"

</details>

<details>
    <summary>
        <h3>Declaração múltipla</h3>
    </summary>

    var a, b int = 1, 2
</details>

## Pacotes + Funções essenciais

##### fmt

<details>
    <summary>
        .Printf() &rarr; Print de texto formatado
    </summary>

    Exemplo: 

    fmt.Printf("Type: %T - Value: %v", true, true)
    Resultado: Type: bool - Value: true

    --

    Notações:
    
    %v -> Printa o valor da variável
    %T -> Printa o tipo da variável
</details>

## Funções

- Se a função tiver um nome iniciado por letra minúscula, ela é privada, ou seja, só pode ser utilizada no próprio pacote
- Se for com letra maíuscula, pode ser usada em outros pacotes também

<details>
    <summary>
        <h3>Declaração</h3>
    </summary>

    // Função obrigatória. A função main sempre é a função principal do Go.
    func main() {
        fmt.Printf("A soma é %v", soma(2,3))
    }

    func soma(x int, y int) int {
        return x + y
    }

</details>

<details>
    <summary>
        <h3>Ignorar retorno de função</h3>
    </summary>

    func main() {
        nome, _ := nomeSobrenome("Gustavo", "Oliveira")

        _, sobrenome = nomeSobrenome("José", "Maria")

        fmt.Println(nome) // Gustavo
        fmt.PrintLn(sobrenome) //  Maria
    }

    func nomeSobrenome(nome, sob string) (string, string) {
        return nome, sobrenome
    }

</details>

OBS: O "_" ignora apenas um parâmetro por vez, ou seja, para cada parâmetro de retorno a ser ignorado, deve-se usar o _

## Listas e dicionários

Existem algumas estruturas para manipular dados:

- Arrays &rarr; Lista com tamanho fixo
- Slices &rarr; Lista com tamanho variável, mais usado que os arrays
- Maps &rarr; Dicionário com tamanho variável

<details>
    <summary>
        <h3>Declarações - Arrays</h3>
    </summary>

    // Declaração explícita
    var a [3]int = [3]int{1, 2, 3}

    // Com valores
    a := [3]int{1, 2, 3}

    // Com inferência
     a := [3]int{1, 2, 3}

     // Automatizada, valores contados pelo compilador
     a := [...]int{10, 20, 30}

</details>

<details>
    <summary>
        <h3>Declarações - Slices</h3>
    </summary>

    // Declaração explícita com valores
    var s []int = []int{1, 2, 3}

    // Com valores
    s := []int{1, 2, 3}

    // Vazio
    var s []int

    // Com make
    s := make([]int, 5)     // Posições zeradas
    s := make([]int, 3, 10) // len=3, cap=10

    // A partir de um array
    a := [5]int{1, 2, 3, 4, 5}
    s := a[1:4]     // slice: [2 3 4]

</details>

<details>
    <summary>
        <h3>Declarações - Maps</h3>
    </summary>

    // Declaração explícita
    var m map[string]int = make(map[string]int)

    // Com valores:
    m := map[string]int{
        "joao": 10,
        "maria": 20,
    }

    // Com inferência
    m := make(map[string]int)

    // Vazio - Evitar usar pois é necessário inicializar o map
    var m map[string]int
    m = make(map[string]int) // Após esse comando é possível utilizar o map

</details>