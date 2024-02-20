Na linguagem Go, `Example` e `Test` são dois tipos diferentes de funções de teste que podem ser usadas para verificar o comportamento do código.

- `Test`: Funções de teste que começam com `Test` são usadas para escrever testes unitários. Eles não retornam nenhum valor e são usados para verificar se uma função ou método está funcionando conforme esperado. Por exemplo, você pode escrever um teste para verificar se uma função retorna o valor correto para um conjunto específico de entradas.

- `Example`: Funções de exemplo que começam com `Example` são usadas para fornecer exemplos de como usar uma função ou método. Esses exemplos são executados como testes, mas o foco principal é demonstrar o uso da função em vez de verificar um comportamento específico. Eles devem ter uma saída documentada que é comparada com a saída real para verificar se o exemplo ainda é válido. Se a saída documentada não corresponder à saída real, o teste falhará.

Aqui está um exemplo de como você pode escrever um `Example` em Go:

```go
package main

import "fmt"

// Exemplo de função que exibe uma saudação.
func ExampleHello() {
    fmt.Println("Olá, mundo!")
    // Output: Olá, mundo!
}
```

Neste exemplo, a função `ExampleHello` imprime "Olá, mundo!" e a saída esperada é documentada abaixo da função. Quando você executa os testes com `go test`, o Go compara a saída da função `ExampleHello` com a saída documentada. Se elas corresponderem, o teste passará; se não, o teste falhará.

Em resumo, `Test` é usado para testar o comportamento de uma função ou método, enquanto `Example` é usado para demonstrar como usar uma função ou método, com uma verificação de saída para garantir que o exemplo ainda seja válido.
