package main

import "fmt"

const limite = 250

func main() {
	suma := 0

	for i := limite; i >= 1; i-- {
		if i%2 == 0 {
			suma += i
		}
	}

	fmt.Println(suma)
}
