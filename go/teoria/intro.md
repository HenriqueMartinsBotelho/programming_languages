Go é bem aprimorada para criação de infraestruturas como servidores web, APIs, sistemas distribuídos, aproveitando recursos de multi-core e rede.

[Meu repositorio de estudos em Go](https://github.com/HenriqueMartinsBotelho/programming_languages/tree/main/go)

### Características

1. É compilada
1. Tem coleta de lixo (garbage collection)
1. Suporte nativo à concorrência com goroutines e channels
1. Estática e fortemente tipada
1. Oferece ferramentas de diagnóstico como o pprof para análise de desempenho.
1. As strings são imutáveis e codificadas em UTF-8
1. As funções são de primeira classe o que significa que podem ser atribuídas a variáveis e passadas como argumentos para outras funções.

Go tem poucos recursos quando comparada a outras linguagens, mas é uma linguagem simples e fácil de aprender.

### Coisas que a linguagem Go não tem

1. Conversões numéricas implícitas
2. Construtores e destrutores
3. Sobrecarga de operadores
4. Valores padrão para parâmetros
5. Herança
6. Exceções
7. Macros
8. Anotações de funções
9. Armazenamento local para threads: Em Go, isso é comumente tratado por meio de goroutines e não threads tradicionais. O armazenamento local para threads, como o "thread-local storage" em algumas outras linguagens, não é uma funcionalidade diretamente oferecida por Go.

---

Threads (ou threads do sistema operacional): Sequências de execução que são agendadas pelo sistema operacional. Elas têm um overhead considerável em termos de memória e tempo de criação/destruição.

Goroutines: Um modelo de concorrência oferecido pelo Go. Goroutines são muito mais leves em termos de recursos quando comparadas às threads tradicionais. Elas são agendadas pelo próprio runtime do Go, não pelo sistema operacional. Isso significa que o Go pode executar milhões de goroutines em uma única máquina.

---

### O sistema de tipos

A linguagem Go apresenta um sistema de tipos projetado para prevenir a maioria dos erros comuns, sem ser excessivamente complexo como em algumas outras linguagens. Seu sistema de tipos é estático, o que significa que o tipo de cada variável é determinado em tempo de compilação. Além disso, Go tem um sistema de tipos forte, impedindo operações inválidas entre tipos diferentes. Vale ressaltar que, enquanto Go oferece um robusto sistema de tipos, os programadores desta linguagem não exploram a tipagem da mesma forma intensa que os programadores de Haskell fazem ao expressar propriedades de segurança e provas baseadas em tipos.

### Hello World e variáveis de ambiente

As variáveis de ambiente são conjuntos de pares chave-valor que configuram o comportamento de alguns programas no sistema. No caso da linguagem de programação Go, elas são essenciais para definir onde o Go e seus componentes estão localizados no sistema.

Para visualizar as variáveis de ambiente do Go, utilize o comando `go env`. Dentre as variáveis listadas, as mais importantes são:

- `GOROOT`: Local onde o Go está instalado.
- `GOPATH`: Local onde os pacotes e projetos Go são armazenados.
- `GOBIN`: Diretório onde os binários Go são instalados.
- `GOENV`: Local onde os arquivos de configuração do Go estão armazenados.
- `GOMODCACHE`: Local onde os módulos Go são armazenados.

---

A variável `GOMODCACHE` é uma variável de ambiente introduzida a partir da versão Go 1.15. Ela indica o diretório base onde os módulos Go baixados (arquivos zip e os arquivos descompactados) são armazenados.

Antes do Go 1.15, esses módulos eram armazenados diretamente no subdiretório `pkg/mod` dentro de `GOPATH`. Com a introdução de `GOMODCACHE`, você pode especificar um diretório diferente para armazenar esses módulos, se desejar.

Se `GOMODCACHE` não estiver definido, o padrão continua sendo o subdiretório `pkg/mod` dentro de `GOPATH`.

Então, em resumo, `GOMODCACHE` aponta para o diretório onde os módulos Go são cacheados no sistema. Isso é especialmente útil quando você está trabalhando com módulos Go e deseja entender ou modificar onde esses módulos são armazenados.

---

Primeiro crie um módulo com o comando `go mod init nome-do-modulo`. Para executar um simples "Hello World" em Go, crie um arquivo chamado `main.go` e adicione o seguinte conteúdo:

```go
package main


func main() {
    Println("Hello World!")
}
```

Para rodar digite `go run main.go`
A partir da versão Go 1.18, uma das adições mais notáveis foi a introdução de um pacote padrão chamado print, que permite aos desenvolvedores imprimir no console sem a necessidade do pacote fmt. Essa mudança foi feita para tornar mais simples e direta a experiência de iniciantes com a linguagem, especialmente para aqueles que estão apenas experimentando ou aprendendo os conceitos básicos.

---

### Entendendo as pastas do diretório GOPATH

O diretório `GOPATH` é o local onde os pacotes e projetos Go são armazenados. Ele é composto por três pastas principais:

1. `bin`: Contém os binários Go.
2. `pkg`: Contém os pacotes Go.
3. `src`: Contém os códigos-fonte Go.

---

### Sobre pacotes

O nome de um pacote Go é o nome da pasta que o contém. A exceção é o pacote `main`, que é o pacote que contém a função `main` e é o ponto de entrada de um programa Go.
Se dois arquivos Go estiverem na mesma pasta, eles pertencem ao mesmo pacote e podem se referenciar diretamente. Se estiverem em pastas diferentes, eles pertencem a pacotes diferentes e devem ser importados para serem referenciados.

### Variáveis e Constantes

- As variáveis são declaradas com a palavra-chave `var` seguida pelo nome da variável e seu tipo.
- As constantes são declaradas com a palavra-chave `const` seguida pelo nome da constante e seu valor.
- Variáveis podem ser declaradas em uma única linha, separadas por vírgulas, e podem ser inicializadas com valores. Se não forem inicializadas, elas recebem o valor padrão do tipo.
- Outra maneira de declarar variáveis é usando `:=` seguida pelo valor da variável. Nesse caso, o tipo da variável é inferido pelo compilador.

```go
var nome string
var (
    nome string = "Henrique"
    idade int
)
b := true
const nome string = "Henrique"
```

### Fontes

- Livro: A linguagem de programação Go
