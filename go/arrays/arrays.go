package arrays

import "fmt"

func Exemplo() {
    // Array
    var array [3]int = [3]int{10, 20, 30}
    fmt.Println("Array:", array)

    // Slice
    slice := []int{10, 20, 30}
    fmt.Println("Slice:", slice)

    // Alterando um elemento
    array[1] = 25
    slice[1] = 25

    // Imprimindo os modificados
    fmt.Println("Modificado Array:", array)
    fmt.Println("Modificado Slice:", slice)

    // A principal diferença é que slices são mais flexíveis e podem ser redimensionados


    //Map
    exemplo := map[string]int{"a": 1, "b": 2}
    fmt.Println("Map:", exemplo)
    fmt.Println("Map['a']:", exemplo["a"])

}
