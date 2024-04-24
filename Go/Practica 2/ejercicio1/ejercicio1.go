package main

import "fmt"

func main() {

	var temps [10]float64

	for i := 0; i < len(temps); i++ {
		fmt.Printf("Ingrese la temperatura del paciente %d: ", i+1)
		fmt.Scanln(&temps[i])
	}

	var alta, normal, baja int
	max := -1.0
	min := 100.0
	for _, temp := range temps {
		if temp > 37.5 {
			alta++
		} else if temp >= 36 && temp <= 37.5 {
			normal++
		} else {
			baja++
		}

		if temp > max {
			max = temp
		}

		if temp < min {
			min = temp
		}
	}

	totalPacientes := alta + normal + baja
	porcentajeAlta := float64(alta) / float64(totalPacientes) * 100
	porcentajeNormal := float64(normal) / float64(totalPacientes) * 100
	porcentajeBaja := float64(baja) / float64(totalPacientes) * 100
	promedio := int((max + min) / 2)

	fmt.Printf("\nPorcentaje de pacientes en cada grupo:\n")
	fmt.Printf("Alta: %.2f%%\n", porcentajeAlta)
	fmt.Printf("Normal: %.2f%%\n", porcentajeNormal)
	fmt.Printf("Baja: %.2f%%\n", porcentajeBaja)

	fmt.Printf("Maximo: %.2f°\n", max)
	fmt.Printf("Minimo: %.2f°\n", min)
	fmt.Printf("Promedio: %.2f°\n", promedio)

}
