package main

import (
	"fmt"
)

func factorial(num int64) int64 {
	if num == 0 {
		return 1
	}
	return factorial(num-1) * num
}

func factorial_it(num int64) int64 {
	r := num
	for {
		num -= 1
		if num == 0 {
			break
		}
		r = r * num
	}

	return r
}

func main() {
	var num int64
	fmt.Scanln(&num)
	fmt.Println(factorial_it(num))
}
