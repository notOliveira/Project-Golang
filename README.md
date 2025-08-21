# Project Go

## [Documentação GO](https://pkg.go.dev/)

## Aprendendo Golang

Não sei muito sobre essa linguagem, então esse repo vai estar meio desorganizado, provavelmente.

#### Para rodar os códigos

```
go run main.go
```

### Para instalar as dependências

```
go mod tidy
```

### Para buildar os códigos

```
go build -o nome_arquivo.exe
```

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

<i></o>OBS: <b>nil</b> é um valor especial que representa a ausência de valor para alguns tipos, como ponteiros, slices, maps, channels e interfaces. Semelhante ao None do Python e null em outras linguagens.</i>

## Variáveis & Constantes

<details>
    <summary>
    <h3>Declaração rápida de variável (inferência de tipo)</h3>
    </summary>

    idade := 30
</details>

<i>OBS: Só pode ser usado dentro de funções. Só funciona para variáveis</i>

### &rarr; &rarr; &rarr; Detalhe ultra mega importante: &larr; &larr; &larr;
Ao declarar funções, caso ela seja iniciada com a letra minúscula, ela será apenas interna ao seu pacote. Do contrário, ela poderá ser exportada e utilizada em outros pacotes do seu projeto Go.

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
        .Sprintf() &rarr; Retorna o valor de uma string formatada
    </summary>

    Exemplo: 
    // Exemplo: 

    fmt.Printf("Type: %T - Value: %v", true, true)
    Resultado: Type: bool - Value: true
    // Resultado: Type: bool - Value: true

    // Para retornar o valor de uma string formatada:

    str := fmt.Sprintf("Número: %d, Texto: %s", 42, "oi")
    fmt.Println(str)
    // Resultado: Número: 42, Texto: oi

    --

    Notações:

    %v -> Printa o valor da variável
    %T -> Printa o tipo da variável
</details>

##### strconv

<details>
    <summary>
        .Itoa() &rarr; Converte int para string
        .Atoi() &rarr; Converte string para int
    </summary>

    // Exemplo:

    str := strconv.Itoa(123) // "123"
    num, err := strconv.Atoi("456") // 456, nil (se a conversão for bem-sucedida)

    if err != nil {
        fmt.Println("Erro na conversão:", err)
    } else {
        fmt.Println("Número convertido:", num)
    }

    // Resultado: Número convertido: 456

    // OBS: A função Atoi retorna dois valores: o número convertido e um erro (se houver). É importante sempre verificar o erro ao fazer a conversão.

    // Outras funções úteis do pacote strconv:
    str := strconv.FormatBool(true) // "true"
    str := strconv.FormatFloat(3.14, 'f', 2, 64) // "3.14"
    str := strconv.FormatInt(12345, 10) // "12345"
    num, err := strconv.ParseBool("true") // true, nil
    num, err := strconv.ParseFloat("3.14", 64) // 3.14, nil
    num, err := strconv.ParseInt("12345", 10, 64) // 12345, nil

    // Sempre verificar o erro ao usar as funções Parse*
    
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

- <b>Arrays</b> &rarr; Lista com tamanho fixo
- <b>Slices</b> &rarr; Lista com tamanho variável, mais usado que os arrays
- <b>Maps</b> &rarr; Dicionário com tamanho variável

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

    // Declaração de mais de um slice ao mesmo tempo
    var s1, s2 []int

</details>

<details>
    <summary>
        <h3>Declarações - Maps</h3>
    </summary>

    // Declaração explícita com inicialização usando make
    var m map[string]int = make(map[string]int)

    // Declaração com inicialização e valores (literal)
    m := map[string]int{
        "joao":  10,
        "maria": 20,
    }

    // Declaração com inferência e inicialização vazia usando make
    m := make(map[string]int)

    // Ou inicialização vazia usando literal
    m = map[string]int{}

    // Declaração sem inicialização (map nil)
    var m map[string]int
    // m está nil, não pode adicionar valores ainda
    // Para usar, precisa inicializar:
    m = make(map[string]int) // Agora pode usar normalmente

</details>

OBS: Os maps no Go não mantêm a ordem dos elementos, ou seja, não é garantido que a ordem de inserção será a mesma na hora de iterar sobre os elementos.

<hr>

Inserindo valores:

```
var s []int
s = append(s, 1, 4, 8) // [1 4 8]

var a [3]int = [3]int{2}
a[2] = 4 // [2 0 4]
```

Para consultar o length:

```
len(slice)
```

## Structs

<details>
    <summary>
        <h3>Declaração</h3>
    </summary>

    type Pessoa struct {
        Nome      string
        Sobrenome string
        Idade     int
    }

    // Declaração de variável do tipo Pessoa
    var gustavo Pessoa

    // Declaração com valores
    gustavo := Pessoa{
        Nome:      "Gustavo",
        Sobrenome: "Oliveira",
        Idade:     30
    }

    É possível também declarar sem os nomes dos campos, mas não é recomendado:
    gustavo := Pessoa{"Gustavo", "Oliveira", 30}

    // Acessando os campos
    fmt.Println(gustavo.Nome) // Gustavo

    // Structs dentro de outras structs
    type Endereco struct {
        Rua    string
        Numero int
    }

    type Pessoa struct {
        Nome      string
        Sobrenome string
        Idade     int
        Endereco  Endereco // Struct dentro de outra struct
    }

    // Acessando os campos da struct interna
    fmt.Println(gustavo.Endereco.Rua) // Rua Exemplo

    // Slice de structs
    pessoas := []Pessoa{
        {Nome: "Gustavo", Sobrenome: "Oliveira", Idade: 30},
        {Nome: "Maria", Sobrenome: "Silva", Idade: 25},
    }

</details>

Podem ser usadas como uma classe, 

## If / Else

<details>
    <summary>
        <h3>Declaração</h3>
    </summary>

    idade := 18

    if idade >= 18 {
        fmt.Println("Maior de idade")
    } else {
        fmt.Println("Menor de idade")
    }

    // Com else if
    if idade < 12 {
        fmt.Println("Criança")
    } else if idade < 18 {
        fmt.Println("Adolescente")
    } else {
        fmt.Println("Adulto")
    }

    // Com declaração curta
    if idade := 18; idade >= 18 {
        fmt.Println("Maior de idade")
    } else {
        fmt.Println("Menor de idade")
    }

    // OBS: A variável idade só existe dentro do if/else

</details>

## Switch / Case

<details>
    <summary>
        <h3>Declaração</h3>
    </summary>

    dia := "Segunda"

    switch dia {
    case "Segunda":
        fmt.Println("Início da semana")
    case "Sábado", "Domingo":
        fmt.Println("Fim de semana")
    default:
        fmt.Println("Dia da semana")
    }

    // Com expressão
    switch {
    case dia == "Segunda":
        fmt.Println("Início da semana")
    case dia == "Sábado" || dia == "Domingo":
        fmt.Println("Fim de semana")
    default:
        fmt.Println("Dia da semana")
    }

    // Com declaração curta
    switch dia := "Segunda"; dia {
    case "Segunda":
        fmt.Println("Início da semana")
    case "Sábado", "Domingo":
        fmt.Println("Fim de semana")
    default:
        fmt.Println("Dia da semana")
    }
</details>

## For (Loops)

<details>
    <summary>
        <h3>Declaração</h3>
    </summary>

    // Loop tradicional
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // Loop com range (para slices, arrays e maps)
    numeros := []int{1, 2, 3, 4, 5}
    for i, num := range numeros {
        fmt.Printf("Índice: %d - Valor: %d\n", i, num)
    }

    // Loop infinito
    for {
        fmt.Println("Loop infinito")
        break // Para sair do loop
        // Se não tiver o break, o loop continuará indefinidamente
        // Para parar o programa, pode usar ctrl + C no terminal
    }  

    // Loop com condição (similar ao while)
    count := 0
    for count < 5 {
        fmt.Println(count)
        count++
    }
</details>  