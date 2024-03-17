package arrays

import "fmt"

/* There are major differences between the ways arrays work in Go and C. In Go,

   - Arrays are values. Assigning one array to another copies all the elements.
   - In particular, if you pass an array to a function, it will receive a copy of the array, not a pointer to it.
   - The size of an array is part of its type. The types [10]int and [20]int are distinct.
*/

func Exemplo() {
	// Array
	var array [3]int = [3]int{10, 20, 30}
	fmt.Println("Array:", array)

	// A principal diferença é que slices são mais flexíveis e podem ser redimensionados
	// Enquando arrays são passados como valor, slices são passados como referência
	slice := []int{10, 20, 30}
	slice = append(slice, 40)
	fmt.Println("Slice:", slice)

	// Alterando um elemento
	array[1] = 25
	slice[1] = 25

	// Imprimindo os modificados
	fmt.Println("Modificado Array:", array)
	fmt.Println("Modificado Slice:", slice)

	//Map
	exemplo := map[string]int{"a": 1, "b": 2}
	fmt.Println("Map:", exemplo)
	fmt.Println("Map['a']:", exemplo["a"])

}
