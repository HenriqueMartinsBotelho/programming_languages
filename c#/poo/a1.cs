using System; 


class A1{
    static void Main(){
        C1 obj1 = new C1();
        obj1.Metodo1();
    }

}

class C1{
    int n1 = 20;

    public void Metodo1(){
        Console.WriteLine("Classe C1");
    }

}