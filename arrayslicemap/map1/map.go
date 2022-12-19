package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[98765432121] = "Pedro"
	aprovados[13243546576] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[13243546576])
	delete(aprovados, 13243546576)
	fmt.Println(aprovados[13243546576])
}