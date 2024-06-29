package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Inicia Goroutine de hello")
	for i := 0; i < 3; i++ {
		fmt.Println(i, " Hello world")
		time.Sleep(3 * time.Second)
	}
	fmt.Println("Termina Goroutine de hello")
}

func main() {
	fmt.Println("Inicia Goroutine del main")
	go hello()
	fmt.Println("Termina Goroutine del main")
}
