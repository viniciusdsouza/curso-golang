package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var prod1 produto
	prod1 = produto{
		nome: "Caneta",
		preco: 1.79,
		desconto: 0.05,
	}
	fmt.Println(prod1, prod1.precoComDesconto())

	prod2 := produto{"Computador", 2500.99, 0.07}
	fmt.Println(prod2, prod2.precoComDesconto())
}
