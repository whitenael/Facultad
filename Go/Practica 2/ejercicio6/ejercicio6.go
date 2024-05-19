package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
)

func sum[T any](a, b []T) []T {
	min := math.Min(float64(len(a)), float64(len(b)))
	rs := make([]T, int(min))

	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		panic("Tipos incompatibles")
	}

	// Creamos un sub-slice vacio para almacenar las sumas de los elementos dependiendo del tipo
	// Debemos asegurarnos que almancera el tipo de numeros correctos
	suma := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(a).Elem()), int(min), int(min))

	for i := 0; i < int(min); i++ {
		valorA := reflect.ValueOf(a[i])
		valorB := reflect.ValueOf(b[i])

		// Sumar los valores
		suma.Index(i).Set(reflect.ValueOf(valorA.Interface().(int) + valorB.Interface().(int)))
	}

	for i := 0; i < int(min); i++ {
		rs[i] = suma.Index(i).Interface().(T)
	}

	return rs
}

func avg(a []int) int {
	var t int
	for _, v := range a {
		t += v
	}

	return t / len(a)
}

func initialize[T any](length int, min, max T) []T {
	result := make([]T, length)

	minT := reflect.TypeOf(min)
	maxT := reflect.TypeOf(max)

	if minT != maxT {
		return nil
	}

	// usamos reflexion para la asignacion de valores random
	// la reflexion nos permite examinar y manipular la estructura interna del programa durante tiempo de ejecucion

	for i := 0; i < length; i++ {
		switch minT.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			minInt := reflect.ValueOf(min).Int()
			maxInt := reflect.ValueOf(max).Int()
			resultValue := reflect.New(reflect.TypeOf(min)).Elem()    // obtenemos el tipo y valor al que apunta la variable min y lo usamos para crear un nuevo tipo y asignarlo a una variable
			resultValue.SetInt(minInt + rand.Int63n(maxInt-minInt+1)) // establecemos el valor de la variable creada usando SetInt para forzarle un valor entero
			result[i] = resultValue.Interface().(T)                   // Tomamos el valor obtenido a traves de reflexion, lo convertimos a tipo generico y los asignamos a la posicion del result
		case reflect.Float32, reflect.Float64:
			minFloat := reflect.ValueOf(min).Float()
			maxFloat := reflect.ValueOf(max).Float()
			resultValue := reflect.New(reflect.TypeOf(min)).Elem()
			resultValue.SetFloat(minFloat + rand.Float64()*maxFloat - minFloat)
			result[i] = resultValue.Interface().(T)
		default:
			panic("Tipo no soportado")
		}

	}

	return result

}

func main() {
	sla := initialize(5, 0, 20)
	slb := initialize(5, 0, 10)

	fmt.Println(sla)
	fmt.Println(slb)

	suma := sum(sla, slb)
	fmt.Println(suma)

}
