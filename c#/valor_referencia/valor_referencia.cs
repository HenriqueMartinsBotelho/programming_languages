// Exemplo passando valor por referência

using System;

class PassandoValorExemplo {
    static void Main() {
        int a = 10;
        int b = 20;

        Console.WriteLine("a = {0}, b = {1}", a, b);

        Troca(ref a, ref b);

        Console.WriteLine("a = {0}, b = {1}", a, b);
    }

    static void Troca(ref int x, ref int y) {
        int temp = x;
        x = y;
        y = temp;
    }
}