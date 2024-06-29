package main

import (
	"fmt"
	"sync"
	"time"
)

func rutina(wg *sync.WaitGroup, id, delay int) {
	defer wg.Done()

	var timeElapsed int
	count := 3

	for i := 0; i < count; i++ {
		fmt.Printf("Rutina: %d \n", id)
		time.Sleep(time.Duration(delay) * time.Second)
		timeElapsed += delay * count
	}
}

func main() {

	// Wait Groups

	var wg sync.WaitGroup

	wg.Add(2)
	go rutina(&wg, 1, 2)
	go rutina(&wg, 2, 2)

	wg.Wait()
	fmt.Println("Done")
	// -------------------------------------------------------

	// Canales

	// Canal sin buffer
	// ch1 := make(chan int)

	// // Ejemplo con canal sin buffer
	// go func() {

	// 	ch1 <- 1                       // Esta línea se bloqueará hasta que alguien reciba del canal ch1
	// 	fmt.Println("Enviado 1 a ch1") // Esta línea se ejecutará después de que el valor se haya recibido
	// }()

	// // Dar tiempo para asegurarse de que la goroutine anterior haya comenzado
	// time.Sleep(time.Second)

	// // Recibir el valor del canal sin buffer
	// val1 := <-ch1
	// fmt.Println("Recibido", val1, "de ch1")

	// ----------------------------------------------------

	// // Ejemplo con canal con buffer

	// Un canal con buffer en Go es un tipo de canal que puede almacenar un número limitado de valores antes de que cualquier goroutine lectora (consumidora)
	// tenga que recibir un valor. Esto se logra especificando la capacidad del canal cuando se crea.

	// Crear un canal con buffer nos permite ALMACENAR valores en el mismo sin que se bloque
	// Un canal se libera una vez que el mensaje/s que se estaba intentando transmitir haya sido recibido por un receptor

	// // Canal con buffer de capacidad 2
	// ch2 := make(chan int, 3)

	// ch2 <- 2
	// fmt.Println("Enviado 2 a ch2") // Esto se imprimirá inmediatamente porque hay espacio en el buffer

	// ch2 <- 3
	// fmt.Println("Enviado 3 a ch2") // Esto también se imprimirá inmediatamente porque hay espacio en el buffer

	// ch2 <- 4
	// fmt.Println("Enviado 4 a ch2") // Esto también se imprimirá inmediatamente porque hay espacio en el buffer

	// val2 := <-ch2
	// fmt.Println("Recibido", val2, "de ch2")

	// val3 := <-ch2
	// fmt.Println("Recibido", val3, "de ch2")

	// val4 := <-ch2
	// fmt.Println("Recibido", val4, "de ch2")

}
