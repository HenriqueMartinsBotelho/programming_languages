/*
Uma analogia interessante para entender os canais em Go seria compará-los a um revezamento de corrida em corredores, onde os atletas passam um bastão de mão em mão. Aqui, os corredores representam as goroutines, que são funções ou métodos executados de forma concorrente, e o bastão representa a informação ou dados que são compartilhados entre eles através dos canais.*/

package concorrencia

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

/*Veja que o código acima não funciona se removermos a palavra go. Ao remover go, a função anônima não é mais lançada como uma goroutine separada. Em vez disso, ela é executada no mesmo fluxo de execução que a função main. Isso introduz um problema: quando a função anônima tenta enviar um valor para o canal (canal <- 10), ela bloqueia até que outro processo esteja pronto para ler do canal. No entanto, como a execução é sequencial e não há outra goroutine para ler do canal neste momento, o programa fica bloqueado indefinidamente esperando por um receptor que nunca virá. Esse tipo de situação é conhecido como deadlock.

Os canais em Go são bloqueantes por padrão, o que significa que a operação de envio (<-) bloqueia até que outro goroutine esteja pronto para receber o valor do canal, e vice-versa. Quando a palavra go é removida, não há concorrência, e portanto, não há outro goroutine ativo para satisfazer a operação de recebimento, resultando em um bloqueio permanente do programa.
*/

// Também é possível usar buffers. Os buffers também bloqueiam. Note que nesse caso o buffer tem tamanho 2 então só podemos
// enviar 2 valores e receber 2 valores.

func buffer_exemplo() {

	canal := make(chan int, 2)
	canal <- 10
	canal <- 20
	fmt.Println(<-canal)
	fmt.Println(<-canal)
}
