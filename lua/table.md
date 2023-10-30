**Tabelas em Lua**

As tabelas são as estruturas de dados mais importantes (e, por concepção, a única estrutura de dados complexas) em Lua, e são a base de todos os tipos de dados possíveis na linguagem. Por meio da utilização de tabelas podemos simular vetores, matrizes, estruturas, objetos, etc.

A tabela é um conjunto de chaves e pares de dados também conhecido como hashed heterogeneous associative array', onde os dados são referenciados por chave. A chave (índice) pode ser de qualquer tipo de dado exceto nulo. Um inteiro com chave valor 1 é considerada distinta da chave de uma string "1".

1. **Criação de Tabelas**:
   Tabelas são criadas usando chaves `{}`. Exemplo:

   ```lua
   vazia = {}
   numeros = {1, 2, 3, 4, 5}
   ```

2. **Acessando e Modificando Tabelas**:
   Você pode acessar e modificar valores em uma tabela usando índices.

   ```lua
   frutas = {"maçã", "banana", "laranja"}
   print(frutas[1])  -- imprimirá "maçã"
   frutas[1] = "uva"  -- substitui "maçã" por "uva"
   ```

3. **Tabelas como Dicionários**:
   Em vez de usar índices numéricos, você pode usar chaves.

   ```lua
   pessoa = {nome="João", idade=30}
   print(pessoa["nome"])  -- imprimirá "João"
   print(pessoa.nome)     -- alternativamente, imprimirá "João"
   ```

4. **Iteração em Tabelas**:
   Lua fornece o loop `pairs` para iterar sobre tabelas.

   ```lua
   for chave, valor in pairs(pessoa) do
       print(chave, valor)
   end
   ```

5. **Tabelas e Referências**:
   Quando atribuímos uma tabela a outra variável ou a passamos como argumento para uma função, estamos passando uma referência para a tabela original, e não uma cópia dela.

   ```lua
   a = {x=1, y=2}
   b = a
   b.x = 5
   print(a.x)  -- imprimirá 5, já que b e a referenciam a mesma tabela
   ```

6. **Metatables e Tabelas**:
   Metatables podem ser associadas a tabelas para alterar ou expandir seu comportamento padrão. Por exemplo, podemos usar metatables para definir operações matemáticas entre tabelas.
   ```lua
   setmetatable(tabelaA, {
       __add = function(tabelaA, tabelaB)
           -- lógica de adição
       end
   })
   ```
