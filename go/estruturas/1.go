package estruturas

import "fmt"

type Endereco struct {
	Logradouro string
	Numero int
	Bairro string
}

type Cliente struct {
	Nome string
	Idade int
	Ativo bool
	Endereco Endereco
}

func Estruturas() {
	henrique := Cliente{Nome: "Henrique", Idade: 21, Ativo: true}
	henrique.Endereco = Endereco{Logradouro: "Rua das Flores", Numero: 123, Bairro: "Centro"}
	fmt.Println(henrique)
}