package concorrencia

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/* Quando você usa apenas a palavra-chave 'go' para iniciar uma goroutine, você está simplesmente dizendo ao Go para executar a função em uma nova thread leve (goroutine) de forma assíncrona. Isso significa que o controle é retornado imediatamente ao ponto de chamada, e a execução do programa continua sem esperar que a goroutine termine. Esta é uma forma poderosa de realizar operações em paralelo, mas não oferece nenhum mecanismo direto para sincronizar estas goroutines ou esperar por suas conclusões. Por exemplo, se o seu programa principal terminar antes de suas goroutines, elas serão interrompidas, e você pode nunca ver seus resultados.

Usando sync.WaitGroup

sync.WaitGroup é uma estrutura fornecida pelo pacote sync que ajuda a sincronizar múltiplas goroutines, permitindo que o programa espere até que um conjunto de goroutines tenha terminado de executar. Aqui estão os componentes chave de como sync.WaitGroup é usado:

Add(int): Incrementa o contador do WaitGroup por um valor inteiro, indicando quantas goroutines precisam ser esperadas para concluir.

Done(): Decrementa o contador do WaitGroup, indicando que uma goroutine concluiu sua execução. Normalmente, essa chamada é deferida ou colocada ao final da função goroutine.

Wait(): Bloqueia a execução até que o contador do WaitGroup chegue a zero, indicando que todas as goroutines registradas com Add chamaram Done.

No exemplo abaixo, wg.Add(2) prepara o WaitGroup para esperar por duas goroutines. Cada goroutine chama wg.Done() quando termina, e wg.Wait() no final de GrupoDeEsperaExemplo bloqueia até que ambas as goroutines tenham chamado Done(). Isso assegura que o programa principal espere pela conclusão de ambas as goroutines antes de prosseguir, proporcionando uma sincronização limpa e evitando a terminação precoce do programa.
var wg sync.WaitGroup
*/

var wg sync.WaitGroup

func GrupoDeEsperaExemplo() {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	wg.Add(2)

	go func1()
	go func2()

	fmt.Println("Goroutines", runtime.NumGoroutine())

	wg.Wait()
}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func1", i)
		time.Sleep(time.Millisecond * 500)
	}
	wg.Done()
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("func2", i)
		time.Sleep(time.Millisecond * 500)
	}
	wg.Done()
}
