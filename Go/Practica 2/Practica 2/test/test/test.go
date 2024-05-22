package main

import "fmt"

// Declaracion de tipos

// Igual al tipo subyacente pero incompatibles entre si
// El casteo es permitido mientras ambos tengan el mismo tipo subyacente

type Celsius float64
type Fahrenheit float64

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// Los parametros, aun siendo punteros seran por valor
// pero no su contenido
func zero(xPtr *int) {
	*xPtr = 0
}

func main() {

	// ----------------------------------------------------------------------- //
	// Ejemplo de uso de punteros

	// var p *int
	// i := 42
	// p = &i
	// fmt.Println(*p) // 42

	// Ejemplos con el alocador de memoria new
	// p := new(int)
	// q := new(int)

	// *p = 10
	// *q = 5

	// fmt.Println(*p, *q) // 10 5
	// fmt.Println(p, q)
	// q = p               // ahora q apuntara a la misma direccion de memoria que p
	// fmt.Println(*p, *q) // 10 10
	// fmt.Println(p, q)

	// ----------------------------------------------------------------------- //
	// Testeo de funcion zero para demostracion de valores pasados por parametro

	// demostracion con variable entera pasando el puntero al valor
	x := 5
	zero(&x)

	fmt.Println(x)

	// demostracion con variable de tipo puntero
	xPtr := new(int)
	zero(xPtr)

	fmt.Println(xPtr)

	// ----------------------------------------------------------------------- //

	// Declaracion de tipos

}
