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

	for i := 0; i < op.Len(); i++ {
		if position == 0 {
			if element == op.data[i].num {
				op.data[i].ocur++
			} else {
				newData := make([]OptimumNode, len(op.data)+1)

				copy(newData, op.data[:i])

				newNode := OptimumNode{element, 1}
				newData[i] = newNode

				copy(newData[i+1:], op.data[i:])

				op.data = newData
			}
			break
		} else {
			min := math.Min(float64(position), float64(op.data[i].ocur))
			position -= int(min)
		}
	}

	return element
}

func (op OptimumSlice) SliceArray() [int]{
	
}

func (op OptimumSlice) ToString() string {
	var str string
	for i := 0; i < len(op.data); i++ {
		str += fmt.Sprintf("%d[%d] ", op.data[i].num, op.data[i].ocur)
	}

	return str
}
