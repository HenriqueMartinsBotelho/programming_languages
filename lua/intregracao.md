A linguagem Lua é frequentemente usada como uma linguagem de script para ser incorporada em programas escritos em outras linguagens. Uma das razões para isso é que a Lua foi projetada para ter uma API simples, porém poderosa, para permitir a integração com código C e C++. Através dessa API, é possível chamar funções Lua a partir de C/C++ e vice-versa.

A integração se dá basicamente através da pilha Lua. Quando você quer passar informações entre C e Lua, você coloca e retira valores dessa pilha.

Aqui está um exemplo básico de como isso se dá:

### Integrando Lua com C

1. **Chamando uma função Lua a partir de C**

Suponha que você tenha o seguinte script Lua, salvo como `script.lua`:

```lua
function somar(a, b)
    return a + b
end
```

Você pode chamar essa função a partir de um programa C assim:

```c
#include <lua.h>
#include <lauxlib.h>
#include <lualib.h>

int main(void) {
    lua_State *L = luaL_newstate();
    luaL_openlibs(L);  // carrega as bibliotecas padrão

    // Carrega e executa o script lua
    if (luaL_dofile(L, "script.lua") != LUA_OK) {
        // Trate o erro aqui
    }

    lua_getglobal(L, "somar");  // pega a função somar
    lua_pushnumber(L, 5);       // empilha o primeiro argumento
    lua_pushnumber(L, 3);       // empilha o segundo argumento

    if (lua_pcall(L, 2, 1, 0) != LUA_OK) {  // chama a função com 2 argumentos e 1 valor de retorno
        // Trate o erro aqui
    }

    double resultado = lua_tonumber(L, -1);  // pega o resultado do topo da pilha
    printf("Resultado: %f\n", resultado);

    lua_close(L);  // fecha o ambiente Lua
    return 0;
}
```

2. **Chamando uma função C a partir de Lua**

Primeiro, você define uma função em C:

```c
static int c_somar(lua_State *L) {
    double a = lua_tonumber(L, 1);
    double b = lua_tonumber(L, 2);
    lua_pushnumber(L, a + b);
    return 1;  // indica que a função retorna um valor
}
```

Em seguida, registre a função no ambiente Lua:

```c
luaL_Reg funcs[] = {
    {"c_somar", c_somar},
    {NULL, NULL}
};

luaL_newlib(L, funcs);
lua_setglobal(L, "meu_modulo");
```

Agora, no script Lua, você pode chamar a função assim:

```lua
local resultado = meu_modulo.c_somar(5, 3)
print(resultado)
```
