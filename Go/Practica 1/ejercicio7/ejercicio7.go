package main

import "fmt"

const temp_alta = 37.5
const temp_baja = 36
const pac_totales = 10

func calcularPorcentaje(x int32) float64 {
	return (float64(x)) / pac_totales * 100
}

func main() {

	var g1, g2, g3 int32
	var act float64

	for i := 0; i < pac_totales; i++ {
		fmt.Printf("Ingrese la temperatura del paciente: ")
		fmt.Scanln(&act)

		switch {
		case act >= temp_alta:
			g1++
		case act <= temp_baja:
			g3++
		case act >= temp_baja && act <= temp_alta:
			g2++
		default:
		}
	}

	fmt.Println("Porcentaje de pacientes del Grupo 1: ", calcularPorcentaje(g1))
	fmt.Println("Porcentaje de pacientes del Grupo 2: ", calcularPorcentaje(g2))
	fmt.Println("Porcentaje de pacientes del Grupo 3: ", calcularPorcentaje(g3))

}
