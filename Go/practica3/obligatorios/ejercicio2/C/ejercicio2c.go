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

func despachador(clientesGlobal chan Cliente, cajas []chan Cliente, longitudes []int, mu *sync.Mutex) {
	for cliente := range clientesGlobal {
		// Se utiliza mutex para proteger el recurso de longitudes
		// De no protegerse se pueden (y se provocaran) condiciones de carrera
		mu.Lock()
		minLen := longitudes[0]
		minIndex := 0
		for i := 1; i < numCajas; i++ {
			if longitudes[i] < minLen {
				minLen = longitudes[i]
				minIndex = i
			}
		}
		longitudes[minIndex]++
		mu.Unlock()

		//
		cajas[minIndex] <- cliente
	}
	for _, caja := range cajas {
		close(caja)
	}
}

func main() {
	clientesGlobal := make(chan Cliente, numClientes)

	cajas := make([]chan Cliente, numCajas)
	// en longitudes guardaremos la cola de cada caja, para luego utilizar ese valor para asignar la cola a cada cliente
	longitudes := make([]int, numCajas)
	var mu sync.Mutex

	for i := range cajas {
		cajas[i] = make(chan Cliente, numClientes)
		longitudes[i] = 0
	}

	var wg sync.WaitGroup
	wg.Add(numClientes)

	for i := 0; i < numCajas; i++ {
		go caja(i+1, cajas[i], &wg)
	}

	go despachador(clientesGlobal, cajas, longitudes, &mu)

	for i := 1; i <= numClientes; i++ {
		clientesGlobal <- Cliente{id: i}
	}
	close(clientesGlobal)

	wg.Wait()

	fmt.Println("Todos los clientes han sido atendidos.")
}
