// Canais servem para transmitir dados entre go routines.
/*
Canais são como corredores em uma corrida de bastão.
Eles tem que passar o bastão de maneira sincronizada.

*/

package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)

	go func() {
		canal <- 10
	}()

	fmt.Println(<-canal)
}
