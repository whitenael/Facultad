package main

import "fmt"

func main() {
	var sl slice
	var opSlice OptimumSlice
	sl = []int{3, 3, 3, 3, 3, 7, 7, 7, 7, 7, 7, 7, 23, 23, 23, 23, 23, 23, 3, 3, 3, 3, 3, 3, 3, 3, 7, 5, 5, 5}

	opSlice = New(sl)
	fmt.Println(opSlice.ToString())

	opSlice.Insert(9, 6)
	fmt.Println(opSlice.ToString())

	opSlice.Insert(9, 6)
	fmt.Println(opSlice.ToString())

	opSlice.Insert(10, 10)
	fmt.Println(opSlice.ToString())

	opSlice.Insert(23, 8)
	fmt.Println(opSlice.ToString())

}
