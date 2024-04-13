package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Mismo funcionamiento que el ejercicio1, solamente que ahora esta parametrizado

func modificarTexto(src, target, replace string) string {

	// chequeamos que se contenga la palabra
	ok := strings.Contains(strings.ToLower(src), strings.ToLower(target))
	ret := src

	if ok == true {
		ret = convertirTexto(src, target, replace)
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

func convertirTexto(src string, target string, replace string) string {
	indexes := obtenerIndicencias(src, target)
	nsrc := src
	var mod string

	if len(indexes) == 0 {
		return src
	}

	for i := 0; i < len(indexes); i++ {
		replace = modificarCapitalizacion(src, target, replace, indexes[i])
		mod = nsrc[:indexes[i]] + replace + nsrc[indexes[i]+len(target):]
		nsrc = mod
	}

	return mod
}

func modificarCapitalizacion(src, target, replace string, index int) string {
	nsrc := src[index : index+len(target)]
	nsrc_rune := []rune(nsrc)
	nreplace := []rune(strings.ToUpper(replace))

	for i, v := range nsrc_rune {
		ok := unicode.IsLower(v)
		if ok == true {
			nreplace[i] = unicode.ToLower(nreplace[i])
		}
	}

	return string(nreplace)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	// comando para ejecutar desde la terminal -> 
	// go run ejercicio2.go < rp.input02.txt
	fmt.Println(modificarTexto(line, "miércoles", "automóvil"))
}
