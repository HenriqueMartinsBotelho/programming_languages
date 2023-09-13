using System;

class ExemploConstrutores {
    static void Main(){
        Animal animal = new Animal(10, "Gato");
        Console.WriteLine(animal.vida);
    }
}

class Animal {
    public int vida;
    public string nome;

    public Animal(int v, string n){
        vida = v;
        nome = n;
    }

    // Método destrutor

    ~Animal(){
        Console.WriteLine("Objeto destruído");
    }

}
