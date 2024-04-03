### O que são Go Routines?

Go routines são funções que são executadas de forma concorrente com outras go routines dentro de um mesmo programa Go. Elas são uma das características fundamentais da linguagem Go, permitindo que os desenvolvedores implementem concorrência de maneira simples e eficaz. Uma go routine é iniciada usando a palavra-chave `go` seguida por uma chamada de função. Isso sinaliza ao Go runtime para executar essa função de forma concorrente.

```go
package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("world")
    say("hello")
}
```

No exemplo acima, a função `say` é chamada duas vezes, uma diretamente e outra como uma go routine. Isso faz com que as duas chamadas sejam executadas concorrentemente, potencialmente em paralelo se o sistema tiver múltiplos núcleos de CPU disponíveis. A saída do programa pode variar. A ausência da palavra-chave go na chamada de say("hello") significa apenas que ela será executada na goroutine principal e não iniciará uma nova goroutine.

### Go Routines vs. Threads

Embora go routines possam parecer threads tradicionais à primeira vista, elas são fundamentalmente diferentes em vários aspectos:

- **Leveza**: Go routines são muito mais leves do que threads. Enquanto uma thread pode exigir uma quantidade significativa de memória overhead por thread (tipicamente 1MB ou mais), uma go routine inicia com uma stack de apenas algumas kilobytes. Isso significa que é possível ter milhares, ou até mesmo milhões, de go routines em um programa Go sem esgotar a memória do sistema.

- **Gerenciamento**: O Go runtime gerencia as go routines internamente, incluindo a multiplexação das go routines em threads do sistema operacional. Isso simplifica o modelo de programação, já que o desenvolvedor não precisa se preocupar diretamente com a criação, a sincronização e o gerenciamento de threads.

- **Cooperatividade**: Enquanto threads tradicionais são preemptivas, go routines operam em um modelo mais cooperativo, com a troca de contexto ocorrendo durante operações de bloqueio ou quando explicitamente cedendo controle ao scheduler do Go usando `runtime.Gosched()`. Isso resulta em um uso mais eficiente dos recursos.

### Exemplo Prático

Vamos considerar um exemplo prático onde queremos fazer múltiplas requisições a uma API de forma concorrente:

```go
package main

import (
    "fmt"
    "net/http"
    "time"
)

func fetchURL(url string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()

    fmt.Printf("%s: %d [%v]\n", url, resp.StatusCode, time.Since(start))
}

func main() {
    urls := []string{
        "http://example.com",
        "http://example.net",
        "http://example.org",
    }

    for _, url := range urls {
        go fetchURL(url)
    }

    // Aguardar um momento antes de finalizar
    time.Sleep(5 * time.Second)
}
```

Neste exemplo, fazemos requisições HTTP concorrentes a três URLs diferentes usando go routines. Isso ilustra como as go routines podem ser usadas para realizar tarefas I/O-bound de forma eficiente.

---

## Canais

Canais são uma maneira de passar dados entre go routines. Isso pode ser usado para sincronizar uma go routine com outra, ou para passar dados de uma go routine para outra.
