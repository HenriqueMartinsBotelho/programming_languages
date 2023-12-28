No exmplo na pasta 'estruturas', a diferença entre as duas declarações de struct em Go está na maneira como o campo `Endereco` é definido e utilizado.

1. Na primeira declaração:

   ```go
   type Cliente struct {
       Nome string
       Idade int
       Ativo bool
       Endereco Endereco
   }
   ```

   Aqui, `Endereco` é definido como um campo nomeado dentro da struct `Cliente`. Isso significa que, ao acessar o endereço de um `Cliente`, você precisa usar o nome do campo, como em `cliente.Endereco`.

2. Na segunda declaração:

   ```go
   type Cliente struct {
       Nome string
       Idade int
       Ativo bool
       Endereco
   }
   ```

   Neste caso, `Endereco` é um campo anônimo (ou incorporado) na struct `Cliente`. Isso permite que você acesse os campos de `Endereco` diretamente através de uma instância de `Cliente`, como se fossem parte da própria struct `Cliente`. Por exemplo, se `Endereco` tem um campo chamado `Rua`, você pode acessá-lo diretamente com `cliente.Rua` em vez de `cliente.Endereco.Rua`.

Ambas as abordagens têm seus usos dependendo do que você deseja alcançar em termos de organização de código e clareza. A incorporação é útil para composição ou para simular herança, enquanto campos nomeados são mais explícitos e podem ser mais claros em certos contextos.
