package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

const (
	producers = 2
	consumers = 2
	count     = 3
)

func producer(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		time.Sleep(time.Duration(rand.Float64() * float64(time.Second)))
		rand := rand.IntN(100)
		fmt.Printf("Produciendo %d... \n", rand)
		c <- rand
	}
}

func consumer(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		item := <-c
		fmt.Printf("Consumiendo %d... \n", item)
	}
}

func main() {

	c := make(chan int, producers*consumers)
	var wg sync.WaitGroup

	for i := 0; i < producers; i++ {
		wg.Add(1)
		go producer(c, &wg)
	}

	for i := 0; i < consumers; i++ {
		wg.Add(1)
		go consumer(c, &wg)
	}

	wg.Wait()
	close(c)

}
