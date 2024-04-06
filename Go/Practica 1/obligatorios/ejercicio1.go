package main

import (
	"fmt"
	"strings"
	"unicode"
)

const str = "MARTES"
const str_replace = "JUEVES"

func procesarTexto(f string) string {

	s := strings.ToUpper(f)                // pasamos la frase a uppercase
	index := strings.Index(s, str_replace) // buscamos el indice donde empieza la palabra a reemplazar
	var ns string                          // creamos una variable para almacenar la frase modificada

	// si el indice es diferente de -1, significa que la frase contiene el str a reemplazar

	if index != -1 {
		ns = convertirTexto(f, index)
	}

	return ns
}

func convertirTexto(s string, index int) string {
	ns := []rune(str) // creamos una runa a partir del texto a reemplazar
	rs := []rune(s)   // creamos una runa de la frase para recorrerla

	// recorremos la frase desde el indice de inicio hasta el largo de la frase a reemplazar

	for i := index; i < len(str); i++ {
		if unicode.IsLower(rs[i]) {
			ns[i] = unicode.ToLower(ns[i])
			rs[i] = ns[i]
		}
	}

	return string(rs)
}

func main() {
	t := "hoy es jueves"
	fmt.Println(procesarTexto(t))
}
