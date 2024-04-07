package main

import (
	"fmt"
	"strings"
)

func modificarTexto(src, target, replace string) string {

	// chequeamos que se contenga la palabra
	ok := strings.Contains(src, target)
	ret := src

	if ok == true {
		ret = convertirTexto(src, target, replace)
	}

	return ret

}

func obtenerIndicencias(phrase, word string) []int {
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

func convertirTexto(src string, target string, replace string) string {
	// Buscamos la palabra a reemplazar en la cadena original
	indexes := obtenerIndicencias(src, target)
	nsrc := src
	var mod string

	// Si no encontramos la palabra, devolvemos la cadena original sin cambios
	if len(indexes) == 0 {
		return src
	}

	// Reemplazamos la palabra en la cadena original
	for i := 0; i < len(indexes); i++ {
		mod = nsrc[:indexes[i]] + replace + nsrc[indexes[i]+len(target):]
		nsrc = mod
	}

	return mod
}

func main() {
	str := "Hoy es miércoles, de nuevo miércoles miércoles"
	fmt.Println(modificarTexto(str, "miércoles", "automóvil"))
}
