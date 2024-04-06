package main

import (
	"fmt"
	"math"
)

func calcularValor(x float64) float64 {
	if x < -18 {
		return math.Abs(x)
	} else if x >= -18 && x <= -1 {
		return math.Mod(x, 4)
	} else if x >= 1 && x < 20 {
		return math.Pow(x, 2)
	} else {
		return x * -1
	}
}

func calcularValorSwitch(x float64) float64 {
	switch {
	case x < -18:
		return math.Abs(x)
	case x >= -18 && x < -1:
		return math.Mod(x, 4)
	case x >= 1 && x <= 20:
		return math.Pow(x, 2)
	default:
		return x * -1
	}
}

func main() {
	var numero float64

	fmt.Print("Ingrese un numero: ")
	fmt.Scanln(&numero)

	valor := calcularValorSwitch(numero)

	fmt.Print(valor)
}
