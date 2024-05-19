package main

import (
	"fmt"
)

// definimos la interfaz
type Collection[T any] interface {
	IsEmpty() bool
	Len() int
	ToString() string
	Iterate(func(int) int) int
}

// definimos el tipo
type Stack[T any] struct {
	data LinkedList[T]
}

// definimos los metodos
func New[T any]() Stack[T] {
	return Stack[T]{}
}

func (self Stack[T]) IsEmpty() bool {
	return self.data.IsEmpty()
}

func (self Stack[T]) Len() int {
	return self.data.Len()
}

func (self Stack[T]) ToString() string {
	return self.data.ToString()
}

func (self *Stack[T]) Iterate(f func(int) int) {
	self.data.Iterate(f)
}

// definimos la pila como una FIFO
func (self *Stack[T]) push(dato T) {
	self.data.PushFront(dato)
}

// definimos el pop de estilo generico
func (self *Stack[T]) pop() (T, bool) {
	if self.IsEmpty() {
		var zero T
		return zero, false
	}

	data := self.data.FrontElement()
	self.data.Remove()
	return data, true
}

func (self *Stack[T]) FrontElement() (T, bool) {
	if self.IsEmpty() {
		var zero T
		return zero, false
	}

	return self.data.FrontElement(), true
}

func main() {
	var stack Stack[any]
	stack.push(1)
	stack.push(2)
	stack.push("c")
	stack.push("d")
	stack.push(false)
	stack.push("f")
	stack.push(7)

	fmt.Println(stack.ToString())

	fmt.Println(stack.pop())

	fmt.Println(stack.ToString())
}
