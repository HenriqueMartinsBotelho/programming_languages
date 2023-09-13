/* 

Type casting is when you assign a value of one data type to another type.

There are two types of casting:

Implicit Casting (automatically) - converting a smaller type to a larger type size
char -> int -> long -> float -> double

Explicit Casting (manually) - converting a larger type to a smaller size type
double -> float -> long -> int -> char


A classe Convert possui métodos para conversão de tipos de dados.



*/

using System;

class Conversoes{
    static void Main(){
        int a = 10;
        double b = a; // automatic type casting

        Console.WriteLine(a);
        Console.WriteLine(b);

        double c = 3.14;
        int d = (int) c; // manual type casting

        Console.WriteLine(c);
        Console.WriteLine(d);

        
        int f = Convert.ToInt32(4.74);
        double g = Convert.ToDouble(6);
        
        Console.WriteLine(f);
        Console.WriteLine(g);

    }
}