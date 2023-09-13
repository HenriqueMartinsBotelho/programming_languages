// C# possui uma instrução chamada goto que permite que você pule para uma seção específica do código.

using System;

class GotoExample{
    static void Main(){ 
        int i = 2;
        if(i == 1)
            goto a;
        else
            goto b;   
        a:
            Console.WriteLine("i = 1");
            return;  
        b:  
            Console.WriteLine("i != 1");
            return;  
              
    }
}