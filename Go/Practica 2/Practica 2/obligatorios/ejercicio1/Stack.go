package main

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
func NewStack[T any]() Stack[T] {
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

// definimos la pila como una FIFO
func (self *Stack[T]) Push(dato T) {
	self.data.PushFront(dato)
}

// definimos el pop de estilo generico
func (self *Stack[T]) Pop() (T, bool) {
	if self.IsEmpty() {
		var zero T
		return zero, false
	}

	data := self.data.FrontElement()
	self.data.Remove()
	return data, true
}

// definimos el poll de estilo generico (es igual al pop, solamente que este no devuelve un booleano)
func (self *Stack[T]) Poll() T {
	if self.IsEmpty() {
		var zero T
		return zero
	}

	data := self.data.FrontElement()
	self.data.Remove()
	return data
}

func (self *Stack[T]) FrontElement() (T, bool) {
	if self.IsEmpty() {
		var zero T
		return zero, false
	}

	return self.data.FrontElement(), true
}

func (self *Stack[T]) Iterate(f func(T)) {
	self.data.Iterate(f)
}
