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

type Cliente struct {
	id int
}

func caja(id int, clientes chan Cliente, wg *sync.WaitGroup) {
	for cliente := range clientes {
		fmt.Printf("Caja %d está atendiendo al cliente %d.\n", id, cliente.id)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Printf("Caja %d ha terminado de atender al cliente %d.\n", id, cliente.id)
		wg.Done()
	}
}

func despachador(clientesGlobal chan Cliente, cajas []chan Cliente) {
	i := 0
	for cliente := range clientesGlobal {
		cajas[i] <- cliente
		i = (i + 1) % numCajas
	}
	for _, caja := range cajas {
		close(caja)
	}
}

func main() {
	// Crear un canal global para los clientes
	clientesGlobal := make(chan Cliente, numClientes)

	// Crear canales individuales para cada caja

	// Creamos un slice vacio que contendra los canales de las cajas
	cajas := make([]chan Cliente, numCajas)

	// Inicializamos cada canal
	for i := range cajas {
		cajas[i] = make(chan Cliente)
	}

	// Crear un wait group para esperar a que todos los clientes sean atendidos
	var wg sync.WaitGroup
	wg.Add(numClientes)

	// Lanzar las goroutines para cada caja
	for i := 0; i < numCajas; i++ {
		go caja(i+1, cajas[i], &wg)
	}

	// Lanzar la goroutine del despachador
	go despachador(clientesGlobal, cajas)

	// Enviar clientes al canal global
	for i := 1; i <= numClientes; i++ {
		clientesGlobal <- Cliente{id: i}
	}

	close(clientesGlobal)

	// Esperar a que todos los clientes sean atendidos
	wg.Wait()

	fmt.Println("Todos los clientes han sido atendidos.")
}
