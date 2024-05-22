package main

import "fmt"

// Definición de tipo enumerativo para días de la semana
type DiaSemana int

const (
	Lunes DiaSemana = iota
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
	Domingo
)

func main() {
	// Uso del tipo enumerativo
	dia := Lunes
	fmt.Println("Hoy es", dia)
}
