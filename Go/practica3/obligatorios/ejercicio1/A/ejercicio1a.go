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

// Implementacion del algoritmo Criba de Eratostenes para calcular numeros primos menores o iguales a un N dado
func numerosPrimos(n int, wg *sync.WaitGroup, c chan []int) {
	// nos aseguramos que cuando termine la go rutine se limpie el wait group
	defer wg.Done()
	var primeNumbers []int

	for i := 0; i <= n; i++ {
		if esPrimo(i) {
			primeNumbers = append(primeNumbers, i)
		}
	}

	// enviar los numeros primos encontrados al canal
	c <- primeNumbers
}

func main() {

	// marcamos tiempo de inicio
	start := time.Now()

	// Convertir el argumento de la línea de comandos a un entero.
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 1 {
		fmt.Println("Por favor, introduzca un número entero positivo.")
		return
	}

	c := make(chan []int)
	var wg sync.WaitGroup
	wg.Add(1)

	// Mandamos a llamar a la rutina (1)
	go numerosPrimos(n, &wg, c) // (1)

	// wg.Wait() bloquea esta goroutine hasta que todas las goroutines hayan llamado a wg.Done(), asegurando que todo el trabajo de verificación de números primos esté completo.
	// close(c) se llama inmediatamente después de wg.Wait() para cerrar el canal, indicando que no habrá más datos enviados a c

	go func() {
		wg.Wait()
		close(c)
	}()

	for primeNumbers := range c {
		fmt.Printf("Números primos menores o iguales a %d: %v\n", n, primeNumbers)
	}

	// marcamos tiempo de final
	elapsed := time.Since(start)
	log.Printf("Tiempo total: %s", elapsed)
}
