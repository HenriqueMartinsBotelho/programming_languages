
# Estruturas e Interfaces em Go

Em Go, as estruturas (structs) são tipos de dados compostos que agrupam variáveis sob um nome. Essas variáveis, conhecidas como campos, podem ter tipos diferentes. Go também oferece o conceito de interfaces, que são um meio de definir um conjunto de métodos assinaturas que um tipo deve implementar.

## Estruturas (Structs)

Go suporta dois tipos de declarações de estruturas, diferenciadas pela forma como os campos, especialmente os estruturais, são incorporados.

### Primeira Declaração: Campo Nomeado

Na primeira forma de declaração, os campos dentro da estrutura são nomeados explicitamente. Isso promove clareza na definição da estrutura e no acesso aos seus campos.

```go
type Cliente struct {
    Nome     string
    Idade    int
    Ativo    bool
    Endereco Endereco
}
```

Neste exemplo, `Endereco` é um campo nomeado dentro da estrutura `Cliente`, exigindo que o acesso a este campo seja feito através do nome do campo, por exemplo, `cliente.Endereco`.

### Segunda Declaração: Campo Anônimo (Incorporado)

A segunda forma permite a incorporação de uma estrutura dentro de outra sem nomear explicitamente o campo. Este método facilita o acesso aos campos da estrutura incorporada diretamente.

```go
type Cliente struct {
    Nome     string
    Idade    int
    Ativo    bool
    Endereco
}
```

Aqui, `Endereco` é um campo anônimo dentro de `Cliente`, permitindo que os campos de `Endereco` sejam acessados como se fossem parte da estrutura `Cliente` diretamente, por exemplo, `cliente.Rua`, simplificando a sintaxe quando comparado ao acesso através de um campo nomeado.

## Interfaces

Interfaces em Go definem um conjunto de métodos que um tipo deve implementar. Elas são fundamentais para a programação polimórfica em Go, permitindo que diferentes tipos sejam tratados de forma uniforme, desde que implementem os mesmos métodos.

```go
type Forma interface {
    Area() float64
    Perimetro() float64
}
```

Neste exemplo, qualquer tipo que tenha os métodos `Area` e `Perimetro` implementados é considerado um tipo `Forma`. Isso permite o uso de tipos diversos de forma intercambiável, desde que cumpram com a assinatura de métodos definida pela interface, facilitando a criação de código genérico e reutilizável.

---

A relação entre estruturas (structs) e interfaces em Go é fundamental para entender como a linguagem gerencia tipos e polimorfismo. Uma estrutura em Go é uma coleção de campos, enquanto uma interface define um conjunto de métodos (sem implementação). Uma interface é satisfeita por qualquer tipo que implemente todos os métodos declarados nela. Vamos detalhar como as estruturas se relacionam com interfaces e esclarecer a questão de como definir que um método deve implementar uma determinada estrutura.

### Relação Entre Estruturas e Interfaces

1. **Implementação Implícita:** Em Go, uma estrutura implementa uma interface implicitamente ao implementar todos os métodos da interface. Não é necessário declarar explicitamente que uma estrutura implementa uma interface.

2. **Polimorfismo:** Interfaces permitem polimorfismo. Uma função que aceita uma interface como argumento pode receber qualquer tipo que implemente essa interface, incluindo estruturas. Isso permite escrever funções genéricas e reutilizáveis.

3. **Desacoplamento:** Interfaces promovem desacoplamento entre a definição de uma operação e sua implementação. Isso significa que você pode mudar a implementação de uma estrutura sem alterar o código que depende da interface que a estrutura implementa.

### Como Definir Que um Método Deve Implementar Uma Determinada Estrutura

Em Go, você não "diz" explicitamente a um método para implementar uma estrutura; em vez disso, você define que um tipo (geralmente uma estrutura) deve implementar uma interface ao codificar os métodos exigidos pela interface. Se uma estrutura possui todos os métodos definidos em uma interface, então ela implicitamente implementa essa interface.

**Exemplo:**

Suponha que temos a seguinte interface:

```go
type Animador interface {
    Animar() string
}
```

E a estrutura `Boneco`:

```go
type Boneco struct {
    Nome string
}
```

Para dizer que `Boneco` deve implementar a interface `Animador`, você simplesmente implementa o método `Animar` para `Boneco`:

```go
func (b Boneco) Animar() string {
    return b.Nome + " está animando!"
}
```

Agora, qualquer função que aceite a interface `Animador` como argumento pode receber uma instância de `Boneco`, porque `Boneco` implementa todos os métodos de `Animador`.

Correto, em Go, não existem as palavras-chave `extends` ou `implements` que são comuns em muitas outras linguagens de programação orientadas a objetos, como Java. Go adota uma abordagem diferente para a herança e implementação de interfaces, baseando-se nos princípios de composição e implementação implícita de interfaces.

### Implementação Implícita de Interfaces

Em Go, uma interface é um tipo que define um conjunto de métodos, mas não implementa esses métodos por si só. Um tipo (geralmente uma estrutura) implementa uma interface simplesmente por implementar seus métodos. Não há necessidade de declarar explicitamente que um tipo implementa uma interface. Essa é a implementação implícita de interfaces.

Por exemplo, se você tem uma interface `Leitor` com um método `Ler()`, qualquer tipo que tenha um método `Ler()` "implementa" a interface `Leitor` automaticamente.

### Composição sobre Herança

Go prefere composição sobre herança. Em vez de criar uma cadeia de tipos derivados (como você faria com `extends` em outras linguagens), em Go, você incorpora um tipo dentro de outro para alcançar reutilização de código e extensão de funcionalidades. Isso é feito usando o que é conhecido como "campos anônimos" ou "embedding" em structs.

Por exemplo, se você tem uma struct `Pessoa` e quer criar uma struct `Aluno` que "herda" todas as propriedades de `Pessoa`, você pode incorporar `Pessoa` em `Aluno` desta forma:

```go
type Pessoa struct {
    Nome string
    Idade int
}

type Aluno struct {
    Pessoa // Incorporação de Pessoa em Aluno
    Matricula string
}
```

Neste caso, `Aluno` tem acesso a todos os campos e métodos de `Pessoa`, além de seus próprios campos e métodos. Isso é feito sem precisar de `extends` ou qualquer outra palavra-chave específica para herança.
