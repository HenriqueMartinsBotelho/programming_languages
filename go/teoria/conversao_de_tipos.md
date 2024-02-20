# Conversão de Tipos na Linguagem Go

A linguagem Go é conhecida por sua tipagem estática, o que significa que os tipos de dados estão vinculados às variáveis e não aos valores. Isso implica que, uma vez que você define uma variável como `int`, ela só poderá ser uma `int` e não poderá atribuir-lhe uma `string` sem converter o tipo de dados da variável [0][3].

## Convertendo Tipos Numéricos

Em Go, existem vários tipos numéricos, incluindo inteiros e números de ponto flutuante. Às vezes, você pode precisar converter entre esses tipos numéricos para otimizar o desempenho ou preservar a precisão das operações matemáticas [0].

### Exemplo de Conversão de Inteiros para Floats

```go
var numInt int =  10
var numFloat float64 = float64(numInt)
fmt.Println(numFloat) // Saída:  10
```

Neste exemplo, convertemos um inteiro para um float usando a conversão de tipo embutida do Go.

### Conversão de Floats para Inteiros

```go
var numFloat float64 =  10.5
var numInt int = int(numFloat)
fmt.Println(numInt) // Saída:  10
```

Aqui, convertemos um float para um inteiro, o que trunca a parte decimal do número.

## Conversão de Strings para Números e vice-versa

Para converter strings para números e vice-versa, você pode usar o pacote `strconv`. Este pacote fornece funções para converter strings em números e números em strings [1].

### Exemplo de Conversão de String para Int

```go
import (
    "fmt"
    "strconv"
)

func main() {
    str := "123"
    num, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(num) // Saída:  123
}
```

Neste exemplo, usamos `strconv.Atoi` para converter uma string em um inteiro.

### Exemplo de Conversão de Int para String

```go
import (
    "fmt"
    "strconv"
)

func main() {
    num :=  123
    str := strconv.Itoa(num)
    fmt.Println(str) // Saída: "123"
}
```

Aqui, usamos `strconv.Itoa` para converter um inteiro em uma string.

## Conclusão

A conversão de tipos em Go é uma habilidade essencial que permite manipular dados de maneiras diferentes. Compreender como converter entre tipos numéricos e strings, por exemplo, é crucial para lidar com entradas do usuário e realizar operações matemáticas. Ao dominar esses conceitos, você estará mais preparado para escrever programas eficientes e robustos em Go [0].
