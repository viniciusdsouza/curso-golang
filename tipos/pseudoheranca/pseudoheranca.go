package main

import "fmt"

type carro struct {
	nome string
	velocidadeAtual int
}

type ferrari struct {
	carro // campos anônimos
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Printf("A ferrari %s está com o turbo ligado? %v\n", f.nome, f.turboLigado)
	fmt.Println(f)
}