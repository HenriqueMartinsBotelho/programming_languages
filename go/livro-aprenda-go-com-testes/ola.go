package main

import "fmt"

const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "
const espanhol = "espanhol"
const frances = "frances"
const portugues = "portugues"


func Ola(nome string, idioma string) string {

	if nome == "" {
		nome = "mundo"
	}

	return prefixodeSaudacao(idioma) + nome
}

// Em nossa assinatura de função, criamos um valor de retorno chamado (prefixo string)
// Isso criará uma variável chamada prefixo na nossa função e lhe será atribuída o valor "zero".
// Isso depende do tipo, por exemplo para int, será 0 e para string será "". 

func prefixodeSaudacao(idioma string) (prefixo string) {

	switch idioma {
	case frances:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	default:
		prefixo = prefixoOlaPortugues
	}

	return
	
}


func main() {
	fmt.Printf(Ola("mundo", "espanhol"))
}