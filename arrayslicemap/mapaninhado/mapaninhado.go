package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64 {
		"G": {
			"Gabriela Silva": 12345.12,
			"Guga Pereira": 24151.24,
		},
		"J": {
			"Joana Carla": 14521.52,
			"Jonas Ferreira": 5351.32,
		},
		"P": {
			"Paola Oliveira": 14526.35,
			"Paulo José": 12525.52,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}