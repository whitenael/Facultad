package main

import (
	"fmt"
	"sync"
)

func sieveOfEratosthenes(n int, wg *sync.WaitGroup, c chan []int) {
	// nos aseguramos que cuando termine la go rutine se limpie el wait group
	defer wg.Done()

	primes := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	// se itera desde 2 hasta la 2 hasta la raiz de n eliminando todos los multiplos de los numeros primos
	// nos aseguramos que p cuadrado sea menor a n, es decir, que p <= raiz de n
	for p := 2; p*p <= n; p++ {
		if primes[p] {
			// marcamos los multiplos de los primos como false
			for i := p * p; i <= n; i += p {
				primes[i] = false
			}
		}
	}

	var primeNumbers []int
	for p := 2; p <= n; p++ {
		if primes[p] {
			primeNumbers = append(primeNumbers, p)
		}
	}

	//fmt.Println(primeNumbers)

	// enviar los numeros primos encontrados al canal
	c <- primeNumbers
}

func main() {
	n := 20

	c := make(chan []int)
	var wg sync.WaitGroup
	wg.Add(3)

	// mandamos a llamar a la rutina (1) (2) (3)
	go sieveOfEratosthenes(10, &wg, c) // (1)
	go sieveOfEratosthenes(20, &wg, c) // (2)
	go sieveOfEratosthenes(30, &wg, c) // (3)

	go func() {
		wg.Wait() // esperamos a que termine de ejecutarse la rutina (1) (2) (3)
		close(c)  // cerramos el canal al finalizar la rutina anonima
	}()

	fmt.Println(c)

	for primeNumbers := range c {
		fmt.Printf("Números primos menores o iguales a %d: %v\n", n, primeNumbers)
	}

}
