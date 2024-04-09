package main

import (
	"fmt"
	"strings"
	"unicode"
)

func modificarTexto(src, target string) string {

	ok := strings.Contains(strings.ToUpper(src), strings.ToUpper(target))
	ret := src

	if ok == true {
		ret = convertirTexto(src, target)
	}

	return ret
}

func obtenerIndicencias(src, target string) []int {
	var indices []int
	lcTarget := strings.ToLower(target)
	lcSrc := strings.ToLower(src)

	index := 0
	for {
		i := strings.Index(lcSrc[index:], lcTarget)
		if i == -1 {
			break
		}
		indices = append(indices, index+i)
		index += i + len(target)
	}

	return indices
}

func convertirTexto(src, target string) string {
	indexes := obtenerIndicencias(src, target)
	nsrc := src
	var mod string

	if len(indexes) == 0 {
		return src
	}

	// Se reutiliza la logica anterior
	// Para este caso, esta metodologia resulta ineficiente, ya que se reemplaza la palabra por una nueva cuando podria cambiarse caracter por caracter a medida que se recorre
	for i := 0; i < len(indexes); i++ {
		target = modificarCapitalizacion(src, target, indexes[i])
		mod = nsrc[:indexes[i]] + target + nsrc[indexes[i]+len(target):]
		nsrc = mod
	}

	return mod
}

func modificarCapitalizacion(src, target string, index int) string {
	nsrc := src[index : index+len(target)]
	nsrc_rune := []rune(nsrc)
	ntarget := []rune(strings.ToUpper(target))

	// ajuste para la modificacion de capitalizacion, ahora en lugar de copiarla, la invierte
	for i, v := range nsrc_rune {
		ok := unicode.IsLower(v)
		// si es minuscula, la pasa a mayuscula, y viceversa
		if ok == true {
			ntarget[i] = unicode.ToUpper(ntarget[i])
		} else {
			ntarget[i] = unicode.ToLower(ntarget[i])
		}
	}

	return string(ntarget)
}

func main() {
	str := "Parece peqUEño, pero no es tan pequeÑo el PEQUEÑO"
	fmt.Println(modificarTexto(str, "PEQUEÑO"))
}
