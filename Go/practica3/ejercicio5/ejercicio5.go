package main

import "fmt"

// Un PRODUCTOR se tiene que encargar de proveer una linea donde los CONSUMIDOR pueda acceder a ella

func prod(c chan int) {
	c <- 3
}

func consum(c chan int) int {
	return <-c
}

func main() {

	c := make(chan int, 2)
	go prod(c)
	fmt.Println(consum(c))

}
