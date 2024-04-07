package main

import (
	"fmt"
)

func main() {
	str := []rune("MIÉRCOLES")
	str2 := "MIÉRCOLES"

	fmt.Println(len(str))
	fmt.Println(len(str2))

	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]))
	}

	fmt.Println("-----------------")

	for i := 0; i < len(str2); i++ {
		fmt.Println(string(str2[i]))
	}
}
