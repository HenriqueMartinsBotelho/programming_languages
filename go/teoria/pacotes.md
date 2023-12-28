## O que são Pacotes em Go?

Os pacotes são coleções de arquivos Go dentro do mesmo diretório.

### Estrutura de um Pacote

```go
package nomeDoPacote

import "outroPacote"

func FuncaoDoPacote() {
    // Implementação da função
}
```

Cada arquivo em um pacote Go começa com a declaração `package`, seguida pelo nome do pacote. O nome do pacote é geralmente o mesmo nome do diretório que contém os arquivos.

## Visibilidade de Funções e Variáveis

Diferente de linguagens como JavaScript, que utilizam palavras-chave como `export` para definir a visibilidade de funções e variáveis, Go adota uma abordagem mais simples. Em Go, se o nome de uma função, variável ou tipo começa com uma letra maiúscula, ele é acessível de outros pacotes (público). Se começa com uma letra minúscula, é acessível apenas dentro do próprio pacote (privado).

### Exemplo de Visibilidade

```go
package pacoteExemplo

// Exportada: acessível fora do pacote
func FuncaoPublica() {}

// Não exportada: acessível apenas dentro do pacote
func funcaoPrivada() {}
```

## Importando Pacotes

Você pode importar pacotes usando a declaração `import`:

```go
import (
    "fmt"
    "os"
    "github.com/user/projeto/pacote"
)
```

## O que são Módulos em Go?

Módulos são uma coleção de pacotes Go que são lançados juntos. Eles representam a unidade mais alta de distribuição e versionamento de código em Go.

### Iniciando um Módulo

Para iniciar um módulo, use o comando:

```bash
go mod init github.com/seunome/seuprojeto
```

Isso cria um novo arquivo `go.mod` no diretório do seu projeto, que define o nome do módulo e suas dependências.

## Gerenciamento de Dependências

O arquivo `go.mod` gerencia as dependências do seu projeto. Quando você importa pacotes de outros módulos, Go atualiza automaticamente esse arquivo.

### Exemplo de `go.mod`

```plaintext
module github.com/seunome/seuprojeto

go 1.16

require (
    github.com/outraLib v1.2.3
)
```
