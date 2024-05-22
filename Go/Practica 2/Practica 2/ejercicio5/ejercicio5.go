package main

import "fmt"

type vector [3]float64

func Initilize(v *vector, f float64) {
	for i := range v {
		v[i] = f
	}
}

func Sum(v1, v2 vector) vector {

	var r vector

	for i := range v1 {
		r[i] = v1[i] + v2[i]
	}

	return r
}

func SumInPlace(v1, v2 *vector) {

	for i := range v1 {
		v1[i] += v2[i]
	}
}

func Print(v vector) {
	for i := range v {
		fmt.Println(v[i])
	}
}

func main() {

	var v1 vector
	var v2 vector

	Initilize(&v1, 4)
	Initilize(&v2, 4)

	//v3 := Sum(v1, v2)

	//Print(v3)

	SumInPlace(&v1, &v2)

	Print(v1)

}
