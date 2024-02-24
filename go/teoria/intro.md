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
1. Não importamos arquivos mas sim pacotes que são diretórios que contém arquivos Go.

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

Em GO é importante saber trabalhar bem com ARQUIVOS e entender profundamente os TIPOS DE DADOS e suas operações.

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

Pacotes Go (abreviadamente pkg) nada mais são do que diretórios no espaço de trabalho Go que contém arquivos de origem Go ou outros pacotes Go.
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

### nil

Em Go, é típico para funções que podem falhar retornar dois valores: o resultado da função (neste caso, o hashedPassword como uma string) e um valor de erro para indicar se a operação foi bem-sucedida ou não.

Quando uma função retorna nil como o valor de erro, isso indica que a função completou sua tarefa sem erros. Se a função retornar um valor de erro diferente de nil, isso indica que algo deu errado.

###

A linguagem de programação Go, também conhecida como Golang, foi projetada pela Google para ser eficiente, legível e, acima de tudo, simples. Go é uma linguagem moderna que combina a eficiência de compilação e a segurança de tipos de linguagens estaticamente tipadas como C e C++ com a facilidade de uso de linguagens dinâmicas como Python e JavaScript. Aqui estão algumas das características de Go em termos de paradigmas de programação:

1. **Fortemente Tipada**: Go é uma linguagem fortemente tipada, o que significa que o tipo de cada variável é conhecido em tempo de compilação. Isso ajuda a evitar muitos erros em tempo de execução relacionados a operações inválidas entre tipos incompatíveis.

2. **Compilada**: Go é uma linguagem compilada, o que significa que o código fonte é transformado em código de máquina que é diretamente executado pelo sistema operacional. Isso contribui para a eficiência e performance do código Go.

3. **Concorrência como Primeira Classe**: Uma das características mais notáveis de Go é seu suporte nativo à concorrência. Go introduz goroutines, que são funções que podem ser executadas simultaneamente, e canais, que permitem a comunicação segura entre goroutines. Isso torna a programação concorrente mais acessível e menos propensa a erros.

4. **Sintaxe Simples e Minimalista**: Go foi projetada para ser simples e fácil de ler. Ela elimina muitos dos construtos desnecessários encontrados em outras linguagens, como parênteses em instruções condicionais (if, for) e pontos e vírgulas no final de cada linha, exceto quando são necessários para separar declarações na mesma linha.

5. **Garbage Collection**: Go possui coleta de lixo automática, o que significa que o programador não precisa gerenciar manualmente a alocação e desalocação de memória. Isso reduz a possibilidade de vazamentos de memória e outros erros relacionados à gestão de memória.

6. **Pacotes e Módulos para Gerenciamento de Dependências**: Go utiliza um sistema de pacotes para organizar o código e um sistema de módulos para gerenciar dependências. Isso facilita a manutenção de grandes bases de código e a colaboração entre desenvolvedores.

7. **Interface e Tipos Embutidos**: Go suporta interfaces e tipos embutidos (embedding) para facilitar a reutilização de código e a implementação de polimorfismo. Isso permite uma forma flexível de composição de objetos sem necessitar de herança.

8. **Ferramentas Integradas**: Go vem com um conjunto rico de ferramentas integradas para formatação de código, teste, e gerenciamento de pacotes, o que melhora a eficiência do desenvolvimento e a qualidade do código.

9. **Sem Herança de Classes**: Ao contrário de muitas outras linguagens orientadas a objetos, Go não suporta herança de classes. Em vez disso, ela promove a composição de interfaces e a reutilização de código através de tipos embutidos.

10. **Tipos e Métodos**: Em Go, tipos podem ter métodos associados. Isso permite que funções sejam chamadas no contexto de um tipo específico, facilitando a organização do código e a implementação de comportamentos específicos de tipos.

Essas características fazem de Go uma linguagem poderosa e flexível, adequada para uma ampla gama de aplicações, desde sistemas de backend e serviços de nuvem até ferramentas de linha de comando e desenvolvimento de sistemas distribuídos.
