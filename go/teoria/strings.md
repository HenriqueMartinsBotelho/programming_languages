Strings em Go são uma sequência de bytes imutáveis, frequentemente usadas para armazenar texto. Uma característica importante das strings em Go é que elas são codificadas em UTF-8, um padrão de codificação de caracteres Unicode que pode representar praticamente qualquer caractere de qualquer língua.

Runes, por outro lado, representam um ponto de código Unicode na linguagem Go. Uma rune é um tipo inteiro de 32 bits (int32) que pode armazenar qualquer ponto de código Unicode, tornando-a perfeita para lidar com caracteres individuais de uma string, especialmente quando esses caracteres são representados por mais de um byte em UTF-8.

### Exemplo 1: Iterando sobre Strings e Runes

Para ilustrar como strings e runes funcionam em Go, vamos ver como iterar sobre os caracteres de uma string, tratando corretamente caracteres que são representados por múltiplos bytes em UTF-8.

```go
package main

import (
    "fmt"
)

func main() {
    frase := "Olá, mundo! 🌍"

    // Iterando sobre bytes
    fmt.Println("Iterando sobre bytes:")
    for i := 0; i < len(frase); i++ {
        fmt.Printf("%x ", frase[i])
    }
    fmt.Println()

    // Iterando sobre runes
    fmt.Println("\nIterando sobre runes:")
    for _, r := range frase {
        fmt.Printf("%x ", r)
    }
    fmt.Println()
}
```

No exemplo acima, ao iterar sobre bytes, estamos simplesmente percorrendo cada byte da string. Isso funciona bem para caracteres ASCII, mas caracteres especiais como 'ç' ou emojis como '🌍' são representados por mais de um byte, o que pode resultar em uma iteração incorreta pelos componentes desses caracteres. Ao iterar sobre runes com `for _, r := range frase`, cada 'r' é uma rune que representa um ponto de código Unicode completo, lidando corretamente com caracteres multibyte.

### Exemplo 2: Comparando Strings e Runes

Outra operação comum é comparar strings ou caracteres individuais (runes).

```go
package main

import "fmt"

func main() {
    s1 := "café"
    s2 := "cafe\u0301" // "e" seguido de um combining acute accent

    // Comparando strings
    fmt.Println("Comparando strings:", s1 == s2)

    // Comparando runes
    r1 := 'é'
    r2 := '\u00e9' // Unicode para 'é'

    fmt.Println("Comparando runes:", r1 == r2)
}
```

Neste exemplo, a comparação de strings `s1` e `s2` resulta em `false` porque, embora representem o mesmo texto visualmente, são codificados de maneira diferente em UTF-8. A comparação de runes `r1` e `r2` retorna `true`, pois ambos representam o mesmo ponto de código Unicode para o caractere 'é'.

Esses exemplos demonstram a importância de entender a diferença entre strings e runes ao manipular texto em Go, especialmente quando se lida com caracteres além do conjunto básico ASCII.
