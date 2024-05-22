package main

type slice []int

type OptimumSlice struct {
	data []OptimumNode
}

type OptimumNode struct {
	num  int
	ocur int
}

func New(s slice) OptimumSlice {
	return OptimumSlice{}
}

func Make(s slice) []OptimumNode {

	var op []OptimumNode

	for i := 1; i < len(s); i++ {

		prev := s[i-1]
		current := s[i]

		// creamos un nodo con el elemento previo
		currentNode := OptimumNode{prev, 1}
		// si el elemento previo es igual al actual, sumamos 1 al currentNode y guardamos
		// sino, cambiamos el currentNode por el nodo actual
		if prev == current {
			currentNode.ocur++
			append(op, currentNode)
		} else {
			
		}

	}

	return op
}
