package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("José", "Pq vc não fala comigo?", 3)
	// fale("Carla", "Só posso falar depois de vc!", 1)

	// go fale("Isabela", "Ei...", 500)
	// go fale("David", "Opa...", 500)

	go fale("Melissa", "Entendi!!", 10)
	fale("Mário", "Parabéns!", 5)
}
