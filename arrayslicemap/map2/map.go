package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João": 11325.45,
		"Gabriel Silva": 15523.36,
		"Pietra Guimarães": 13425.34,
	}
	
	funcsESalarios["Rafaela Marques"] = 13142.25
	delete(funcsESalarios, "Invalid")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}