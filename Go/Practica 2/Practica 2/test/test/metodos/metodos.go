package main

import (
	"fmt"
)

// En Go, los métodos son funciones asociadas a un tipo específico, similar a los métodos en otros lenguajes de programación orientados a objetos.
// Sin embargo, Go no tiene clases en el sentido tradicional, en su lugar, utiliza tipos y métodos para lograr la encapsulación y la abstracción.

// Definición del tipo Rectangulo
type Rectangulo struct {
	base   float64
	altura float64
}

// Método Area para calcular el área del rectángulo
func (r Rectangulo) Area() float64 {
	return r.base * r.altura
}

// Método Escalar para escalar el rectángulo
func (r *Rectangulo) Escalar(factor float64) {
	r.base *= factor
	r.altura *= factor
}

func main() {
	// Crear una instancia de Rectangulo
	rect := Rectangulo{base: 10, altura: 5}

	// Llamar al método Area() en la instancia rect
	area := rect.Area()

	fmt.Println("El área del rectángulo es:", area)

	// Se define un tipo Rectangulo con dos campos: base y altura.
	// Se define un método Area() que toma un valor de tipo Rectangulo como receptor (denotado por (r Rectangulo)). Este método calcula y devuelve el área del rectángulo.
	// En la función main(), se crea una instancia de Rectangulo llamada rect con una base de 10 y una altura de 5.
	// Se llama al método Area() en la instancia rect, y el resultado se almacena en la variable area, que luego se imprime.

	fmt.Println("Rectángulo antes de escalar:", rect)

	// Escalar el rectángulo
	rect.Escalar(2)

	fmt.Println("Rectángulo después de escalar:", rect)

	// 	Se imprime el rectángulo antes de escalar.
	// 	Se llama al método Escalar() en la instancia rect, con un factor de escala de 2.
	// 	Se imprime el rectángulo después de escalar, mostrando los cambios en la base y la altura.
}
