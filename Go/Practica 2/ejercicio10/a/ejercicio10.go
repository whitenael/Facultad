package main

import "fmt"

// definimos la interfaz
type Collection[T any] interface {
	IsEmpty() bool
	Len() int
	ToString() string
	Iterate(func(int) int) int
}

// definimos el tipo
type Stack[T any] struct {
	data []T
}

// definimos los metodos

func New[T any]() Stack[T] {
	return Stack[T]{}
}

func (self Stack[T]) IsEmpty() bool {
	return (len(self.data) == 0)
}

func (self Stack[T]) Len() int {
	return len(self.data)
}

func (self Stack[T]) ToString() string {
	var str string
	if self.IsEmpty() {
		return str
	}

	for i := 0; i < self.Len(); i++ {
		str += fmt.Sprint("", self.data[i]) + "-> "
	}

	return str
}

func (s *Stack[T]) Iterate(f func(int) int) {
	if s.IsEmpty() {
		return
	}

	for i := 0; i < s.Len(); i++ {

	}
}

// definimos la pila como una FIFO
func (self *Stack[T]) push(dato T) {
	self.data = append(self.data, dato)
}

// definimos el pop de estilo generico
func (self *Stack[T]) pop() (T, bool) {
	if self.IsEmpty() {
		var zero T
		return zero, false
	}

	val := self.data[len(self.data)-1]
	self.data = self.data[:len(self.data)-1]
	return val, true
}

func (self *Stack[T]) FrontElement() (T, bool) {
	if self.IsEmpty() {
		var zero T
		return zero, false
	}

	val := self.data[0]
	self.data = self.data[1:self.Len()]
	return val, true
}
