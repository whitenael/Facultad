package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

// Función que determina si un número es primo
func esPrimo(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// funcion que comprueba numeros en un rango y envia los primos encontrados a un canal
func checkPrimesInRange(start, end int, wg *sync.WaitGroup, c chan []int) {
	defer wg.Done()

	fmt.Printf("Calculado para el rango: %d-%d: \n", start, end)
	var primes []int
	for num := start; num <= end; num++ {
		if esPrimo(num) {
			primes = append(primes, num)
		}
	}
	c <- primes
}

func main() {

	// marcamos tiempo de inicio
	start := time.Now()

	// convertir el argumento de la línea de comandos a un entero
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 1 {
		fmt.Println("Por favor, introduzca un número entero positivo.")
		return
	}

	// Calculo de cuantas goroutines deberemos usar dependiendo de la entrada N
	numGoroutines := n / 10000
	if numGoroutines < 2 {
		numGoroutines = 2
	} else if numGoroutines > 100 {
		numGoroutines = 100
	}

	rangeSize := n / numGoroutines

	c := make(chan []int)
	var wg sync.WaitGroup

	// Dividir el trabajo entre las goroutines
	for i := 0; i < numGoroutines; i++ {
		start := i*rangeSize + 1
		end := (i + 1) * rangeSize

		// chequeamos el final
		if i == numGoroutines-1 {
			end = n
		}

		// creamos la goroutine con el rango calculado
		wg.Add(1)
		go checkPrimesInRange(start, end, &wg, c)
	}

	// esperamos a que terminen todas las goroutine
	go func() {
		wg.Wait()
		close(c)
	}()

	// Recopilar y mostrar los resultados
	var allPrimes []int
	for primes := range c {
		allPrimes = append(allPrimes, primes...)
	}

	fmt.Printf("Números primos menores o iguales a %d: %v\n", n, allPrimes)

	fmt.Printf("Cantidad de Goroutines usadas: %d \n", numGoroutines)

	// marcamos tiempo de final
	elapsed := time.Since(start)
	log.Printf("Tiempo total: %s", elapsed)
}
