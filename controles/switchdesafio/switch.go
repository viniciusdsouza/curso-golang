package main

import "fmt"

func notaParaConceito(n float64) string {
	switch {
	case n <= 10 && n >= 9:
		return "A"
	case n <= 8 && n >= 7:
		return "B"
	case n <= 6 && n >= 5:
		return "C"
	case n <= 4 && n >= 3:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(7.4))
	fmt.Println(notaParaConceito(9.4))
	fmt.Println(notaParaConceito(5.4))
	fmt.Println(notaParaConceito(3.4))
	fmt.Println(notaParaConceito(1.4))
}
