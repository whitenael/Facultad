package main

import "fmt"

// Las interfaces en Go proporcionan una forma de especificar el comportamiento de un objeto: qué métodos debe implementar un tipo
// para satisfacer esa interfaz. En lugar de una declaración explícita por parte del programador, en Go, una interfaz es implícita y
// cualquier tipo que implemente los métodos requeridos satisface automáticamente la interfaz. Esto permite el polimorfismo en Go sin
// necesidad de herencia explícita.

// Definición de interfaz: Una interfaz en Go es un conjunto de métodos.
// Se define mediante una lista de métodos con sus firmas, pero sin proporcionar la implementación real. Por ejemplo:

type Forma interface {
	Area() float64
	Perimetro() float64
}

// Esta interfaz Forma define que cualquier tipo que tenga métodos Area() y Perimetro() que devuelvan valores de punto flotante puede ser considerado como una forma.

type Rectangulo struct {
	base   float64
	altura float64
}

// Implementación implícita: En Go, no hay una declaración explícita de que un tipo implementa una interfaz.
// Si un tipo tiene los métodos definidos en una interfaz, Go considera automáticamente que ese tipo satisface la interfaz.

// Implementación de los métodos de la interfaz Forma para Rectangulo
// Rectangulo satisface implícitamente la interfaz Forma porque implementa los métodos Area() y Perimetro()

func (r Rectangulo) Area() float64 {
	return r.base * r.altura
}

func (r Rectangulo) Perimetro() float64 {
	return 2*r.base + 2*r.altura
}

// Las interfaces se utilizan principalmente para abstraer el comportamiento de los tipos. Por ejemplo,
// en una función que necesita calcular el área y el perímetro de una forma, podemos usar la interfaz Forma en lugar de tipos concretos como Rectangulo, Circulo, etc.

func ImprimirInformacion(f Forma) {
	fmt.Println("Área:", f.Area())
	fmt.Println("Perímetro:", f.Perimetro())
}

func main() {
	rect := Rectangulo{base: 10, altura: 5}
	ImprimirInformacion(rect) // Rectangulo se pasa como argumento a la función que espera una interfaz Forma

	// ImprimirInformacion acepta cualquier tipo que satisfaga la interfaz Forma.
	// Como Rectangulo implementa los métodos requeridos por Forma, se puede pasar como argumento a ImprimirInformacion.
}
