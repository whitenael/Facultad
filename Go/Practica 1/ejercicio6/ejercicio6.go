package main

import "fmt"

func dividirEnteros(x float64, y float64) float64 {
	max := max(x, y)
	min := min(x, y)

	if min == 0 {
		fmt.Println("No se puede dividir por cero!")
		return 0
	}

	return max / min
}

func main() {

	var num1, num2 float64

	fmt.Println("Ingrese el primer numero")
	fmt.Scanln(&num1)

	fmt.Println("Ingrese el segundo numero")
	fmt.Scanln(&num2)

	fmt.Println(dividirEnteros(num1, num2))
}
