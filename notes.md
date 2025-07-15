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

