package main

import (
	"fmt"
	"strings"
	"unicode"
)

func modificarTexto(src, target string) string {

	// chequeamos que se contenga la palabra
	ok := strings.Contains(strings.ToUpper(src), strings.ToUpper(target))
	ret := src

	if ok == true {
		ret = convertirTexto(src, target)
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

func convertirTexto(src, target string) string {
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
		// cambiamos la capitalizacion del replace
		target = modificarCapitalizacion(src, target, indexes[i])
		mod = nsrc[:indexes[i]] + target + nsrc[indexes[i]+len(target):]
		nsrc = mod
	}

	return mod
}

func modificarCapitalizacion(src, target string, index int) string {
	nsrc := src[index : index+len(target)]     // creamos un nuevo string de la palabra target
	nsrc_rune := []rune(nsrc)                  // creamos runas para recorrer todo, esto se debe ya que las runas y los strings asignan espacio a los caracteres unicode de forma diferente
	ntarget := []rune(strings.ToUpper(target)) // pasamos el replace a mayusculas para formatear todo a partir de minusculas

	for i, v := range nsrc_rune {
		ok := unicode.IsLower(v)
		if ok == true {
			ntarget[i] = unicode.ToUpper(ntarget[i])
		} else {
			ntarget[i] = unicode.ToLower(ntarget[i])
		}
	}

	return string(ntarget)
}

func main() {
	//str := "qqqqqÁ miéRcoLes sfÉsgíó~ñdfdhdhh MiÉRcolEs cgdgdg maRTes miéRcOLÉs miéRcOLEssdsafssfs  .... MMiérCOLES jj"
	str := "Parece peqUEño, pero no es tan pequeÑo el PEQUEÑO"
	fmt.Println(modificarTexto(str, "PEQUEÑO"))
	//fmt.Println(obtenerIndicencias(str, "miércoles"))
	//fmt.Println(modificarCapitalizacion(str, "miércoles", "automóvil", 39))
}
