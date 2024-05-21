package main

import "fmt"

type Date struct {
	dia  int
	mes  int
	anio int
}

func newDate(dia, mes, anio int) Date {
	return Date{dia, mes, anio}
}

type Alumnos struct {
	apellido        string
	nombre          string
	origen          string
	fecha           Date
	presento_titulo bool
	codigo_carrera  string
}

type AlumnosDTO struct {
	apellido string
	nombre   string
}

func Adapt(a Alumnos) AlumnosDTO {
	return AlumnosDTO{a.nombre, a.apellido}
}

func ObtenerMaxKeyMap[T comparable](m map[T]int) (T, int) {
	maxCount := 0
	var max T
	for e, count := range m {
		if count > maxCount {
			max, maxCount = e, count
		}
	}

	return max, maxCount
}

func Resolver(ingresantes LinkedList[Alumnos]) ([]AlumnosDTO, map[int]int, map[string]int, LinkedList[Alumnos]) {

	barilocheIngresantes := []AlumnosDTO{}
	yearCount := make(map[int]int)
	carreraCount := make(map[string]int)
	var tempList LinkedList[Alumnos]

	ingresantes.Iterate(func(ingresante Alumnos) {
		// Sumamos a un contador los ingresantes de bariloche
		if ingresante.origen == "Bariloche" {
			barilocheIngresantes = append(barilocheIngresantes, Adapt(ingresante))
		}

		// Usamos un mapa para llevar el conteo de cada anio y su cantidad de inscriptos
		yearCount[ingresante.fecha.anio]++

		// Repetimos el proceso
		carreraCount[ingresante.codigo_carrera]++

		// Agregamos a una lista temporal solo los alumnos que presentaron el titulo y retornamos
		if ingresante.presento_titulo {
			tempList.PushBack(ingresante)
		}
	})

	return barilocheIngresantes, yearCount, carreraCount, tempList
}

func main() {
	ingresantes := NewList[Alumnos]()

	ingresantes.PushBack(Alumnos{"Perez", "Juan", "Bariloche", newDate(10, 5, 2000), true, "APU"})
	ingresantes.PushBack(Alumnos{"Gomez", "Ana", "Cordoba", newDate(15, 8, 2001), false, "LI"})
	ingresantes.PushBack(Alumnos{"Lopez", "Luis", "Bariloche", newDate(20, 12, 1999), true, "LS"})
	ingresantes.PushBack(Alumnos{"Rodriguez", "Maria", "Buenos Aires", newDate(25, 3, 2000), true, "APU"})

	barilocheIngresantes := []AlumnosDTO{}
	yearCount := make(map[int]int)
	carreraCount := make(map[string]int)
	var tempList LinkedList[Alumnos]

	barilocheIngresantes, yearCount, carreraCount, tempList = Resolver(ingresantes)

	// Informar resultados

	// a) Nombres y apellidos de los ingresantes cuya ciudad origen es “Bariloche”
	fmt.Println("Ingresantes de Bariloche:")
	for _, e := range barilocheIngresantes {
		fmt.Println(e.nombre + ", " + e.apellido)
	}

	// b) Año en que más ingresantes nacieron
	maxYear, maxCount := ObtenerMaxKeyMap(yearCount)
	fmt.Printf("El año con más ingresantes es: %d con %d ingresantes\n", maxYear, maxCount)

	// c) Carrera con la mayor cantidad de inscriptos
	maxCarrera, maxInscriptos := ObtenerMaxKeyMap(carreraCount)
	fmt.Printf("La carrera con mayor cantidad de inscriptos es: %s con %d inscriptos\n", maxCarrera, maxInscriptos)

	// Actualizar la lista de ingresantes
	ingresantes = tempList
	fmt.Println("Lista de ingresantes después de eliminar los que no presentaron título:")
	ingresantes.Iterate(func(ingresante Alumnos) {
		fmt.Println(ingresante.nombre, ingresante.apellido)
	})
}
