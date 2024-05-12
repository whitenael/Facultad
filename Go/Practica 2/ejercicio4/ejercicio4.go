package main

import (
	"fmt"
	"math"
)

const n = 3

func sumatoria(n int) int {
	var result int

	for i := 1; i < n; i++ {
		result += 1 / i
	}

	return result
}

func productoria(n int) int {
	var result int

	for i := 1; i < n; i++ {
		result *= int(math.Pow(float64(i), 3))
	}

	return result
}

func maxmin() {

}

func main() {
	var r [n]int

	sum := sumatoria(n)
	fmt.Println(sum)
	prd := productoria(n)
	fmt.Println(prd)
	r[0] = sum - prd

	fmt.Println(r)

}
