package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numClientes = 10
	numCajas    = 3
)

func cliente(id int, cajas chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulamos una espera antes de ir a la caja
	fmt.Printf("Cliente %d está en la cola.\n", id)
	time.Sleep(time.Second)

	// Se intenta usar una caja disponible, mientras haya cajas en el buffer del canal se procedera, si el buffer se vacia, se bloquea la atencion
	caja := <-cajas
	fmt.Printf("Cliente %d está siendo atendido en la caja %d.\n", id, caja)         // Se abre la caja
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)                    // Simulamos el tiempo de atencion
	fmt.Printf("Cliente %d ha terminado de ser atendido en la caja %d.\n", id, caja) // Se cierra la caja

	// Devolvemos la caja al canal, asi el mismo no se bloquea
	cajas <- caja
}

func main() {

	// Usamos un canal con 3 de buffer para representar las cajas disponibles
	cajas := make(chan int, numCajas)
	for i := 1; i <= numCajas; i++ {
		cajas <- i
	}

	// Creamos un wait group para esperar a que todos los clientes terminen
	var wg sync.WaitGroup
	wg.Add(numClientes)

	// Creamos los clientes y los lanzamos como goroutines
	for i := 1; i <= numClientes; i++ {
		go cliente(i, cajas, &wg)
	}

	// Esperamos a que todos los clientes terminen de ser atendidos
	wg.Wait()

	close(cajas)
	fmt.Println("Todos los clientes han sido atendidos.")
}
