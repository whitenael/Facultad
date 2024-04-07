package main

import (
	"fmt"
	"strings"
	"unicode"
)

const str = "AUTOMÓVIL"
const str_replace = "MIÉRCOLES"

func procesarTexto(f string) string {

	indexes := encontrarOcurrencias(str_replace, f) // buscamos el indice donde empieza la palabra a reemplazar
	var ns string = f                               // creamos una variable para almacenar la frase modificada

	// si el indice es diferente de -1, significa que la frase contiene el str a reemplazar

	if len(indexes) > 0 {
		for i := 0; i < len(indexes); i++ {
			ns = convertirTexto(f, indexes[i], []rune(str))
			fmt.Println(indexes[i])
		}
	}

	return ns
}

func encontrarOcurrencias(word, phrase string) []int {
	var indices []int
	lowerCaseWord := strings.ToLower(word)
	lowerCasePhrase := strings.ToLower(phrase)

	index := 0
	for {
		i := strings.Index(lowerCasePhrase[index:], lowerCaseWord)
		if i == -1 {
			break
		}
		indices = append(indices, index+i)
		index += i + len(word)
	}

	return indices
}

func convertirTexto(s string, index int, ns []rune) string {
	rs := []rune(s) // creamos una runa de la frase para recorrerla

	// recorremos la frase desde el indice de inicio hasta el largo de la frase a reemplazar
	for i := index; i < index+len(ns)-2; i++ {
		if unicode.IsLower(rs[i]) {
			ns[i-index] = unicode.ToLower(ns[i-index])
		}
		rs[i] = ns[i-index]
	}

	return string(rs)
}

func main() {
	t := "hoy es miércoles de miércoles"
	fmt.Println(procesarTexto(t))
}
