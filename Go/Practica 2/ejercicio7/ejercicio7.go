package main

import (
	"fmt"
	"unicode"
)

func decode(src string) {
	nsrc := []rune(src)

	var upCase int
	var loCase int
	var special int
	var numbers int

	for _, v := range nsrc {
		if unicode.IsLetter(v) {
			if unicode.IsUpper(v) {
				upCase++
			} else {
				loCase++
			}
		}

		if unicode.IsNumber(v) {
			numbers++
		}

		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			special++
		}
	}

	fmt.Println(upCase)
	fmt.Println(loCase)
	fmt.Println(special)
	fmt.Println(numbers)
}

func main() {

	char_sqc := "ljDdkhj3453unsd234ms#412908#@!+u3MQ###"
	decode(char_sqc)

}
