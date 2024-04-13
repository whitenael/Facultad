package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Creamos constantes con los valores a buscar y reemplazar
const target = "jueves"
const replace = "martes"

// funcion que evalua si el texto puede cambiarse, de poderse, retorna el texto modificado, de no, retorna el texto sin modificar
func modificarTexto(src string) string {

	// chequeamos que se contenga la palabra
	ok := strings.Contains(strings.ToLower(src), strings.ToLower(target)) // comparamos en minusculas
	// asignamos el valor de retorno
	ret := src

	if ok == true {
		// de haber encontrado al menos una incidencia de la palabra a buscar (target), convertimos el texto
		ret = convertirTexto(src)
	}

	return ret

}

// funcion para obtener los indices de todas las incidencias del target
func obtenerIndicencias(src string) []int {
	var indices []int
	lcTarget := strings.ToLower(target) // comparamos los valores en minusculas
	lcSrc := strings.ToLower(src)

	index := 0
	for {
		// buscamos dentro de un substring variable nuestro target
		// iniciamos desde 0 al final, y a medida que vamos encontrando incidencias, usamos el indice de la ultima incidencia para generar otro substring
		i := strings.Index(lcSrc[index:], lcTarget)

		// en caso de ser -1, rompemos el ciclo y devolvemos los valores encontrados
		if i == -1 {
			break
		}

		// en caso de encontrar un valor, agregamos el indice al arreglo
		indices = append(indices, index+i)
		// y movemos el index actual para no volver a considerarlo dentro del substring
		index += i + len(target)
	}

	return indices
}

func convertirTexto(src string) string {
	// obtenemos los indices del target
	indexes := obtenerIndicencias(src)
	// creamos la variable donde iremos guardando las modificaciones del src
	nsrc := src
	// variables auxiliares
	var mod string
	var rep string = replace

	// nos aseguramos que hayamos encontrado incidencias (metodo redudante, ya que sabemos que las va a contener debido al Contains previo, pero sirve para evitar errores inesperados)
	if len(indexes) == 0 {
		return src
	}

	// Reemplazamos la palabra en la cadena original
	for i := 0; i < len(indexes); i++ {
		// cambiamos la capitalizacion del replace segun como se haya encontrado al target
		rep = modificarCapitalizacion(src, indexes[i])

		// partimos la frase en 3 partes:
		// la primera que va desde el inicio hasta el inicio de la incidencia
		// la segunda sera la indicencia modificada por la palabra replace (con la capitalizacion modificada segun corresponda)
		// y la tercera sera el final de la palabra target hasta el final de la frase (se deberia agregar espacio extra si la palabra replace es mas larga que el target)
		mod = nsrc[:indexes[i]] + rep + nsrc[indexes[i]+len(target):]
		nsrc = mod
	}

	return mod
}

func modificarCapitalizacion(src string, index int) string {
	nsrc := src[index : index+len(target)]       // creamos un nuevo string de la palabra target
	nsrc_rune := []rune(nsrc)                    // creamos runas para recorrer todo, esto se debe ya que las runas y los strings asignan espacio a los caracteres unicode de forma diferente
	nreplace := []rune(strings.ToUpper(replace)) // pasamos el replace a mayusculas para formatear todo a partir de minusculas

	for i, v := range nsrc_rune {
		// si encontramos una letra minuscula, modificamos nuestra palabra replace para hacer el reemplazo
		ok := unicode.IsLower(v)
		if ok == true {
			nreplace[i] = unicode.ToLower(nreplace[i])
		}
	}

	// devolvemos no sin antes volver todo a string
	return string(nreplace)
}

func main() {
	//str := "qqqqqÁ miéRcoLes sfÉsgíó~ñdfdhdhh MiÉRcolEs cgdgdg maRTes miéRcOLÉs miéRcOLEssdsafssfs  .... MMiérCOLES jj"
	str := "hoy es JueVes, de nuevo es jueVes, no soporto los jueveS"
	fmt.Println(modificarTexto(str))
}
