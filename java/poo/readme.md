# Abstração

1. Métodos abstratos não possuem corpo, ou seja, não possuem implementação.
2. Classes abstratas não podem ser instanciadas.
3. Abstração serve para definir um modelo genérico para um grupo de objetos.

## Exemplo

```java

public abstract class Animal {
    public abstract void emitirSom();
}

public class Cachorro extends Animal {
    @Override
    public void emitirSom() {
        System.out.println("Au au");
    }
}

public class Gato extends Animal {
    @Override
    public void emitirSom() {
        System.out.println("Miau");
    }
}

public class Main {
    public static void main(String[] args) {
        Animal animal = new Cachorro();
        animal.emitirSom();
    }
}

```

# Encapsulamento

É definido como o agrupamento de dados em uma única unidade. É o mecanismo que une o código e os dados que ele manipula. Outra forma de pensar no encapsulamento é que ele é um escudo protetor que impede que os dados sejam acessados pelo código fora deste escudo.

## Exemplo

```java

public class Pessoa {
    private String nome;
    private int idade;

    public Pessoa(String nome, int idade) {
        this.nome = nome;
        this.idade = idade;
    }

    public String getNome() {
        return nome;
    }

    public int getIdade() {
        return idade;
    }
}
```
