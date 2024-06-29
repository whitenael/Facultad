package main

import "fmt"

type Map[K comparable, V any] map[K]V

func NewMap[K comparable, V any]() map[K]V {
	return make(map[K]V)
}

func main() {

	m := NewMap[string, int]()
	m["One"] = 1
	m["Two"] = 2
	m["Three"] = 3

	fmt.Println(m)

	n := NewMap[int, string]()
	n[1] = "One"
	n[2] = "Two"
	n[3] = "Three"

	fmt.Println(n)

}
