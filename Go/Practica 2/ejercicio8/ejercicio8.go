package main

import (
	"fmt"
	"strconv"
)

func decimalToBinary(decimal int64, base int) string {
	if 2 < base && base > 36 {
		return "base no valida"
	}
	return strconv.FormatInt(decimal, base)
}

func main() {
	fmt.Println(decimalToBinary(123412, 40))
}
