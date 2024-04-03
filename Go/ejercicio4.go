package main

import "fmt"

func main() {
	limite := 250
	suma := 0

	for i := 1; i <= limite; i++ {
		if i%2 == 0 {
			suma += i
		}
	}

	fmt.Println(suma)
}
