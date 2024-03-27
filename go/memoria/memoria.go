package memoria

import "fmt"

// Neste exemplo, new(int) aloca memória para um inteiro,
//inicializa-o a 0 (o valor zero para inteiros) e retorna um ponteiro para a memória alocada.

func Aprendendo_alocacao() {
	ptr := new(int)
	print(*ptr)
}

// Diferentemente de new, a função make é usada exclusivamente para alocar e inicializar
// objetos de tipos embutidos como slices, maps e channels, que precisam de uma inicialização
// não apenas da memória, mas também de suas estruturas internas.
//make retorna um objeto inicializado (não um ponteiro) do tipo especificado, pronto para uso.

func Aprendendo_make() {
	s := make([]int, 10)
	fmt.Println(s) // Saída: [0 0 0 0 0 0 0 0 0 0]
}
