package main

import "fmt"

func main() {
	limite := 250
	suma := 0

	for i := limite; i >= 1; i-- {
		if i%2 == 0 {
			suma += i
		}
	}

	fmt.Println(suma)
}
