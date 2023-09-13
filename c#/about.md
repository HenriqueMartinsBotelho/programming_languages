# Características da linguagem C#

1. Multiparadigma:

C# é uma linguagem que suporta diferentes estilos de programação, permitindo ao programador utilizar diferentes abordagens dependendo das necessidades do projeto.

2. Tipagem Forte:

Variáveis declaradas em C# são associadas a um tipo específico (como int, string, etc.), e essa associação não pode ser alterada implicitamente.
Ajuda na prevenção de erros, uma vez que o compilador verifica os tipos em tempo de compilação.
Origens e Influências:

3. Sintaxe orientada a objetos derivada do C++.

Incorpora características de outras linguagens, como Object Pascal e Java.

4. Orientação a Objetos:

C# é fortemente orientado a objetos, o que significa que ela permite a criação e manipulação de objetos e classes.
Suporta conceitos como herança, polimorfismo, encapsulamento e abstração. Tudo é um objeto: **System.Object** é a classe base de todo o sistema de tipos de C#;

5. Compilação e Execução:

Código fonte é compilado para Common Intermediate Language (CIL).
CIL é um formato intermediário, agnóstico de linguagem e de plataforma.
A máquina virtual Common Language Runtime (CLR) interpreta ou compila Just-In-Time o CIL para executar o código.

6. Plataforma .NET Framework:

C# foi projetado para ser uma das linguagens primárias da plataforma .NET.
Funciona na Common Language Infrastructure (CLI) que facilita a interoperabilidade entre várias linguagens que compõem o .NET.

# História da linguagem C#

1. Origens e Propósito:

- Parte essencial da plataforma .NET.
- Criada como uma linguagem simples, robusta, orientada a objetos, fortemente tipada e escalável.
- Objetivo de permitir que aplicações sejam executadas em diversos dispositivos, desde PCs até dispositivos móveis.

2. Contexto Histórico:

- A evolução das ferramentas de programação e dispositivos inteligentes trouxe incompatibilidades e novas demandas.
- A Microsoft viu a necessidade de software acessível e disponível em qualquer dispositivo, levando à iniciativa .NET e à criação do C#.

3. Desenvolvimento da Plataforma .NET:

- Bibliotecas originalmente escritas em Simple Managed C (SMC) com um compilador próprio.
- Em janeiro de 1999, Anders Hejlsberg foi recrutado pela Microsoft para liderar o desenvolvimento da nova linguagem.
- Originalmente chamada de "Cool", foi renomeada para "C#" e apresentada ao público em 2000 na Professional Developers Conference (PDC).

4. Principal Contribuidor:

- Anders Hejlsberg, Distinguished Engineer na Microsoft.
- Anteriormente, trabalhou em projetos famosos como Turbo Pascal e Delphi na Borland.

5. Padronização:

- Microsoft submeteu o C# à ECMA para padronização.
- Em dezembro de 2001, a ECMA liberou a especificação ECMA-334 Especificação da Linguagem C#.
- Em 2003, C# tornou-se um padrão ISO (ISO/IEC 23270).

6. Implementações:

- Mono: Implementação open source da Novell.
- dotGNU e Portable.NET: Implementações da Free Software Foundation.
- BDS 2008: Implementação da CodeGear.

## POO

Ao contrário do C++, o C# não suporta herança múltipla (que ocorre quando uma classe é derivada de mais de uma classe base direta).

As classes podem ser abstratas e exigirem que sejam derivadas através da palavra-chave abstract. Nesse tipo de classe métodos não precisam ser obrigatoriamente implementados deixando para suas filhas implementar.

| Modificador de Visibilidade | Descrição                                                                                                                      |
| --------------------------- | ------------------------------------------------------------------------------------------------------------------------------ |
| public                      | Acessível de qualquer outro código em qualquer parte do código ou em qualquer assembly.                                        |
| private                     | Acessível apenas dentro da classe ou estrutura em que é declarado.                                                             |
| protected                   | Acessível apenas dentro da sua classe e por uma classe derivada dessa classe.                                                  |
| internal                    | Acessível de qualquer código no mesmo assembly, mas não de outro assembly.                                                     |
| protected internal          | Acessível dentro de seu assembly e a partir de uma classe derivada. Combina `protected` e `internal`.                          |
| private protected           | Acessível apenas dentro da sua classe ou estrutura, e de classes derivadas no mesmo assembly. Combina `private` e `protected`. |

## Tipos de dados

C# tem um sistema de tipo de dados unificado. Este sistema de tipo unificado é chamado Common Type System (CTS).

Um sistema de tipo unificado implica que todos os tipos, incluindo primitivos, como inteiros, são subclasses da classe System.Object. Por exemplo, cada tipo herda um método ToString().

O CTS separa os tipos de dados em duas categorias:

### Tipos de valor

As instâncias de tipos de valores não têm identidade referencial nem semântica de comparação referencial - comparações de igualdade e desigualdade para tipos de valor comparam os valores de dados reais dentro das instâncias, a menos que os operadores correspondentes estejam sobrecarregados. Tipos de valor são derivados de System.ValueType, sempre têm um valor padrão e sempre podem ser criados e copiados. Algumas outras limitações em tipos de valor são que eles não podem derivar uns dos outros (mas podem implementar interfaces) e não podem ter um construtor padrão explícito (sem parâmetros).

Os tipos de valor normalmente representam dados simples, como valores int ou bool. Os tipos de valor próprios da linguagem são os tipos integrais, o tipo decimal e os tipos de ponto flutuante float (um número de ponto flutuante IEEE de 32 bits) e double, e ainda char (uma unidade de código Unicode de 16 bits) e System.DateTime que identifica um ponto específico no tempo com nanossegundo de precisão). Outros exemplos são enum (enumerações) e struct (estruturas definidas pelo usuário)

### Tipos de referência

Em contraste, os tipos de referência têm a noção de identidade referencial - cada instância de um tipo de referência é inerentemente distinta de todas as outras instâncias, mesmo se os dados dentro de ambas as instâncias forem iguais. Isso é refletido em comparações de igualdade e desigualdade padrão para tipos de referência, que testam a igualdade referencial em vez da estrutural, a menos que os operadores correspondentes estejam sobrecarregados (como o caso para System.String). Em geral, nem sempre é possível criar uma instância de um tipo de referência, nem copiar uma instância existente ou executar uma comparação de valor em duas instâncias existentes, embora tipos de referência específicos possam fornecer tais serviços expondo um construtor público ou implementando um Interface correspondente (como ICloneable ou IComparable). Exemplos de tipos de referência são objeto (a última classe base para todas as outras classes C #), System.String (uma seqüência de caracteres Unicode) e System.Array (uma classe base para todos os arrays C #).
