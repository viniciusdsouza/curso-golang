package main

import (
	"fmt"
	"time"
)

// Channel é a forma de comunicação entre goroutines
// channel é um tipo dentro do Golang

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o channel

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)
	
	a, b := <-c, <-c // recebendo dados do channel
	fmt.Println(a, b)

	fmt.Println(<-c)
}