package main

import (
	"fmt"
)

func main() {
	nome := "Paulo"
	versao := 1.1
	fmt.Println("Olá senhor", nome, "esse programa está na versão", versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O programa escolhido foi: ", comando)
}
