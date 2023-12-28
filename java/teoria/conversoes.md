Converter tipos de dados em Java é uma tarefa comum e existem várias maneiras de fazer isso, dependendo dos tipos de dados envolvidos. Abaixo estão alguns exemplos mostrando diferentes métodos de conversão de tipos em Java:

### 1. Conversão entre Tipos Primitivos

Java suporta conversão automática (upcasting) e explícita (downcasting) entre tipos primitivos.

#### a. Conversão Automática (Upcasting)

```java
int meuInt = 100;
long meuLong = meuInt;  // Conversão automática de int para long
double meuDouble = meuInt;  // Conversão automática de int para double
```

#### b. Conversão Explícita (Downcasting)

```java
double meuDouble = 9.78;
int meuInt = (int) meuDouble;  // Conversão explícita de double para int
```

### 2. Conversão de Strings para Tipos Primitivos

Usando métodos de classes wrapper como `Integer`, `Double`, etc.

```java
String valorString = "123";
int valorInt = Integer.parseInt(valorString);  // Converte String para int
double valorDouble = Double.parseDouble(valorString);  // Converte String para double
```

### 3. Conversão de Tipos Primitivos para Strings

Usando o método `String.valueOf()` ou concatenação.

```java
int meuInt = 200;
String valorString = String.valueOf(meuInt);  // Converte int para String
String valorString2 = meuInt + "";  // Outra maneira de converter int para String
```

### 4. Conversão entre Tipos de Objetos (Casting)

Quando as classes têm uma relação hierárquica.

```java
Object obj = "Uma String";
String str = (String) obj;  // Downcasting de Object para String
```

### 5. Conversão com Classes Wrapper

Usando métodos de classes wrapper para converter entre diferentes tipos primitivos.

```java
Integer meuInteger = 100;
double valorDouble = meuInteger.doubleValue();  // Converte Integer para double
```

### 6. Conversão de Arrays de Tipos Primitivos para Arrays de Objetos

```java
int[] arrayInt = {1, 2, 3};
Integer[] arrayInteger = Arrays.stream(arrayInt).boxed().toArray(Integer[]::new);  // Converte int[] para Integer[]
```

Cada um desses métodos tem suas aplicações específicas e é importante escolher o método correto para evitar erros de tempo de execução e problemas de precisão de dados.
