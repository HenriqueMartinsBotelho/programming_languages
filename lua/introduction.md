1. **Convenções Lexicais**: Identificadores são sequências utilizadas para nomear variáveis, funções ou campos em Lua. Eles podem começar com uma letra ou sublinhado, seguido por letras, números ou sublinhados. Exemplos incluem: `minhaVariavel`, `funcao_1` e `_nome`.

2. **Valores e Tipos**: Lua é uma linguagem dinamicamente tipada. Isso significa que uma variável pode conter diferentes tipos de valores durante sua vida. Exemplos são:

   ```lua
   a = 5          -- a é um número
   a = "olá"      -- agora a é uma string
   ```

3. **Variáveis**: Representam locais de armazenamento que contêm valores. Em Lua, temos:

   ```lua
   local x = 10       -- variável local
   y = 20             -- variável global
   tabela = {}        -- criação de uma tabela
   tabela.campo = 30  -- campo de tabela
   ```

4. **Tabelas**: Uma das características mais distintas de Lua é sua estrutura de dados de tabela. Tabelas são contêineres gerais que podem funcionar como arrays, dicionários, listas e muito mais. Por exemplo:

   ```lua
   frutas = {"maçã", "banana", "laranja"}
   print(frutas[1])   -- imprimirá "maçã"
   ```

5. **Declarações**: São as instruções que o programa realiza. Incluem estruturas de controle, como:

   ```lua
   if x > 10 then
       print("maior")
   else
       print("menor ou igual")
   end
   ```

6. **Expressões**: São combinações que retornam valores. Incluem operações aritméticas, comparações e chamadas de função:

   ```lua
   resultado = x + y * 2 - 3/4
   ```

7. **Regras de Visibilidade**: Variáveis locais têm escopo limitado à função ou bloco onde são declaradas.

   ```lua
   function f()
       local z = 10   -- z só existe dentro de f()
   end
   ```

8. **Tratamento de Erros**: Em Lua, erros podem ser capturados e tratados, permitindo respostas mais graciosas a situações inesperadas.

9. **Coleta de Lixo**: Lua possui um coletor de lixo automático. Isso significa que a memória usada por objetos não referenciados é liberada automaticamente.

10. **Closures**: São funções que capturam variáveis locais de seu escopo envolvente. Permitem construir fábricas de funções, entre outros padrões avançados.

11. **Metatables**: Oferecem mecanismos para personalizar o comportamento de tabelas. Por exemplo, você pode definir como duas tabelas são somadas ou o que acontece ao acessar uma chave inexistente:
    ```lua
    -- Exemplo de metatables para somar tabelas
    setmetatable(tabelaA, {
        __add = function(tabelaA, tabelaB)
            -- define a lógica de adição aqui
        end
    })
    ```
