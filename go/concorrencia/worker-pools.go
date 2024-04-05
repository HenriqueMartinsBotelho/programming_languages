// Worker Pools em programação referem-se a um padrão de design
// //utilizado para gerenciar e limitar o número de threads ou processos que estão executando tarefas em paralelo.
// A ideia básica por trás dos Worker Pools é criar um conjunto pré-definido de "trabalhadores" (threads ou processos), que podem ser reutilizados para executar várias tarefas. Quando uma nova tarefa precisa ser realizada, em vez de criar uma nova thread ou processo, uma tarefa é atribuída a um trabalhador disponível no pool. Se todos os trabalhadores estiverem ocupados, a tarefa pode ser colocada em uma fila até que um trabalhador fique disponível.

package concorrencia

import (
	"fmt"
	"time"
)

// Worker Pool é um padrão de design que pode ser implementado usando goroutines e canais em Go.
// Este padrão é usado para controlar o número de tarefas que são executadas concorrentemente,
// otimizando o uso de recursos e melhorando a eficiência do programa.
// Mesmo em sistemas single-core, um worker pool pode ajudar a melhorar o throughput (vazão) de tarefas que não utilizam a CPU intensivamente, mas que podem se beneficiar da execução paralela devido a esperas em operações de I/O. Isso se alinha ao modelo de concorrência cooperativa de Go, onde goroutines cedem controle voluntariamente ao bloquear em operações de I/O, permitindo que outras tarefas sejam executadas.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func workerPoolExample() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
