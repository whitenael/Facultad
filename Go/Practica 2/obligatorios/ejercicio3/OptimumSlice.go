package main

import (
	"fmt"
	"math"
)

type slice []int

type OptimumSlice struct {
	data []OptimumNode
}

type OptimumNode struct {
	num  int
	ocur int
}

func New(s slice) OptimumSlice {
	return OptimumSlice{Make(s)}
}

func Make(s slice) []OptimumNode {

	if len(s) == 0 {
		panic("No se puede procesar un slice vacio.")
	}

	var op []OptimumNode

	// Inicializamos el primer node
	initialNode := OptimumNode{s[0], 0}
	op = append(op, initialNode)

	for i := 0; i < len(s); i++ {
		lastNode := op[len(op)-1]
		current := s[i]

		if current == lastNode.num {
			op[len(op)-1].ocur++
		} else {
			newNode := OptimumNode{current, 1}
			op = append(op, newNode)
		}
	}

	return op
}

func (op OptimumSlice) IsEmpty() bool {
	return len(op.data) == 0
}

func (op OptimumSlice) Len() int {
	return len(op.data)
}

func (op OptimumSlice) FrontElement() int {

	if op.IsEmpty() {
		return -1
	}

	return op.data[0].num
}

func (op OptimumSlice) LastElement() int {

	if op.IsEmpty() {
		return -1
	}

	return op.data[op.Len()-1].num
}

func (op *OptimumSlice) Insert(element, position int) int {
	if op.IsEmpty() {
		return -1
	}

	var min float64
	aux := position

	if position == 0 {
		newData := make([]OptimumNode, len(op.data)+1)

		newNode := OptimumNode{element, 1}
		newData[0] = newNode

		copy(newData[1:], op.data[:])

		op.data = newData

		return element
	}

	for i := 0; i < op.Len(); i++ {

		min = math.Min(float64(aux), float64(op.data[i].ocur))
		aux -= int(min)

		if aux == 0 {
			if element == op.data[i].num {
				op.data[i].ocur++
			} else {
				diff := math.Abs(float64(op.data[i].ocur) - float64(min))

				if diff > 0 {
					newData := make([]OptimumNode, len(op.data)+2)

					copy(newData, op.data[:i])

					// insertamos el vecino izquierdo
					ln := OptimumNode{op.data[i].num, int(min)}
					newData[i] = ln

					// insertamos el nodo
					newNode := OptimumNode{element, 1}
					newData[i+1] = newNode

					// insertamos el vecino derecho
					rn := OptimumNode{op.data[i].num, int(diff)}
					newData[i+2] = rn

					copy(newData[i+3:], op.data[i+1:])

					op.data = newData
				} else {
					newData := make([]OptimumNode, len(op.data)+1)

					copy(newData, op.data[:i+1])

					newNode := OptimumNode{element, 1}
					newData[i+1] = newNode

					copy(newData[i+2:], op.data[i+1:])

					op.data = newData
				}
			}
			break
		}
	}

	return element
}

func (op OptimumSlice) SliceArray() []int {
	if op.IsEmpty() {
		panic("No se puede procesar un slice vacio.")
	}

	var arr []int

	for i := 0; i < op.Len(); i++ {
		for j := 0; j < op.data[i].ocur; j++ {
			arr = append(arr, op.data[i].num)
		}
	}

	return arr
}

func (op OptimumSlice) ToString() string {
	var str string
	for i := 0; i < len(op.data); i++ {
		str += fmt.Sprintf("%d[%d] ", op.data[i].num, op.data[i].ocur)
	}

	return str
}
