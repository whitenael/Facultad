package main

import "fmt"

// Las funciones anonimas son funciones definidas sin un nombre explicito
// Se pueden utilizar como nombres de primera clase, lo que significa que se pueden asignar a variables, pasar como argumentos a otras funciones y retornar como valores

func Iterate(v []int, f func(int) int) {
	for i := 0; i < len(v); i++ {
		fmt.Println(f(v[i]))
	}
}

func crearMultiplicador(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func mult(n int) int {
	return n * 3
}

func main() {

	// Ejemplo 1:
	saludar := func() {
		fmt.Println("Halo")
	}

	saludar()

	// Ejemplo 2:

	sumar := func(a int, b int) int {
		return a + b
	}

	resultado := sumar(3, 4)
	fmt.Println(resultado) // Imprime 7

	// Funciones Anónimas Como Argumentos
	// Las funciones anonimas como parametros son especialmente utiles para encapsular funcionamiento del codigo y mantener cada funcion separada
	// en un scope, como en la siguiente funcion de operar:

	var datos []int
	datos = append(datos, 10)
	datos = append(datos, 20)
	datos = append(datos, 30)

	Iterate(datos, mult)

	// Funciones Anónimas Como Valores de Retorno
	// Las funciones anónimas también pueden ser retornadas por otras funciones:

	// En este ejemplo, crearMultiplicador es una función que retorna una función anónima.
	// La función anónima multiplica su argumento por un factor dado. La función crearMultiplicador(2) retorna una función que multiplica por dos
	// y esta se asigna a la variable multiplicarPorDos.

	multiplicarPorDos := crearMultiplicador(2)
	fmt.Println(multiplicarPorDos(3))

	// Clausuras
	// Las funciones anónimas en Go también pueden capturar y cerrar sobre variables definidas en su ámbito exterior. Esto se conoce como una clausura.

	// En este ejemplo, la función anónima incrementar captura la variable contador y la modifica cada vez que se llama.
	// Las funciones anónimas tienen acceso a las variables del entorno en el que se crearon, permitiendo este tipo de comportamiento.

	contador := 0
	incrementar := func() int {
		contador++
		return contador
	}

	fmt.Println(incrementar())
	fmt.Println(incrementar())

}
