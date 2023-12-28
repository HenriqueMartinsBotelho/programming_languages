## Tipos de Função e Sintaxe Básica

Em Go, uma função é definida usando a palavra-chave `func`, seguida pelo nome da função, lista de parâmetros (se houver), e o tipo de retorno. Um ponto interessante em Go é a possibilidade de omitir o tipo de parâmetro quando parâmetros consecutivos compartilham o mesmo tipo.

### Exemplo Básico

```go
func soma(a, b int) int {
    return a + b
}
```

Neste exemplo, `a` e `b` são ambos do tipo `int`, portanto, mencionamos o tipo uma única vez.

## Retornando Múltiplos Valores

Go permite que as funções retornem múltiplos valores. Isso é extremamente útil, especialmente ao lidar com funções que podem resultar em erro.

### Exemplo de Múltiplos Retornos

```go
func divisao(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("divisão por zero")
    }
    return a / b, nil
}
```

Aqui, `divisao` retorna um `int` e um `error`. Se `b` for zero, ela retorna um erro.

## O Tipo `error` e `nil`

Em Go, `error` é um tipo embutido usado para expressar erro em operações. Se uma função é bem-sucedida, o erro é tipicamente retornado como `nil`, indicando a ausência de erro.

## Funções Variádicas

Go suporta funções variádicas, que podem aceitar um número variável de argumentos. Isso é feito usando `...` antes do tipo do último parâmetro.

### Exemplo de Função Variádica

```go
func somaVariadica(numeros ...int) int {
    total := 0
    for _, num := range numeros {
        total += num
    }
    return total
}
```

A função `somaVariadica` pode aceitar qualquer número de argumentos `int`.

## Conclusão

Go oferece um sistema de funções poderoso e flexível, com suporte para múltiplos tipos de retorno, funções variádicas e um robusto sistema de tratamento de erros. Essas características tornam Go uma linguagem prática e eficiente para uma ampla gama de aplicações.
