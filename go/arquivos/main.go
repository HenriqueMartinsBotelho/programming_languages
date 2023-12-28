package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// Criando Arquivo
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// Escrevendo String
	tamanho, err := f.WriteString("Olá mundo!")
	if err != nil {
		panic(err)
	}
	fmt.Println(tamanho, "bytes escritos com sucesso")

	// Escrevendo Bytes
	bytes := []byte{65, 66, 67, 68, 69}
	tamanho, err = f.Write(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Println(tamanho, "bytes escritos com sucesso")

	// Fechando e reabrindo o arquivo para leitura
	if err = f.Close(); err != nil {
		panic(err)
	}

	// Lendo Arquivo
	arquivoBytes, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Arquivo:", string(arquivoBytes))

	// Lendo Arquivo por pedaços
	arquivo, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = arquivo.Close(); err != nil {
			panic(err)
		}
	}()

	reader := bufio.NewReader(arquivo)
	buffer := make([]byte, 4)
	for {
		tamanho, err := reader.Read(buffer)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
		fmt.Println(string(buffer[:tamanho]))
	}
}
