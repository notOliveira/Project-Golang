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

## Variáveis

##### Declaração rápida

<code>
    idade := 30
</code>

OBS: Só pode ser usado dentro de funções.

##### Declaração explícita

<code>
    var idade int = 30
</code>

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

