package main

import "fmt"

type Node[T any] struct {
	data int
	next *Node[T]
}

type List[T any] struct {
	head *Node[T]
}

func New[T any]() List[T] {
	return List[T]{}
}

func IsEmpty[T any](self List[T]) bool {
	return self.head == nil
}

func Len[T any](self List[T]) int {
	length := 0
	current := self.head
	for current != nil {
		length++
		current = current.next
	}

	return length
}

func FrontElement[T any](self List[T]) int {
	return self.head.data
}

func Next[T any](self List[T]) List[T] {
	return List[T]{self.head.next}
}

func ToString[T any](self List[T]) string {
	str := ""
	current := self.head
	for current != nil {
		str += fmt.Sprintf("%d", current.data) + "-> "
		current = current.next
	}

	return str
}

func PushFront[T any](self *List[T], elem int) {
	node := &Node[T]{data: elem, next: self.head}
	self.head = node
}

func PushBack[T any](self *List[T], elem int) {
	node := &Node[T]{data: elem, next: nil}
	if self.head == nil {
		self.head = node
		return
	}

	current := self.head
	for current.next != nil {
		current = current.next
	}

	current.next = node
}

func Remove[T any](self *List[T]) int {
	if self.head == nil {
		return 0
	}

	data := self.head.data
	self.head = self.head.next
	return data
}

func Iterate[T any](self List[T], f func(int) int) {
	current := self.head
	for current != nil {
		current.data = f(current.data)
		current = current.next
	}
}

func main() {

	list := List[int]{}

	fmt.Println(ToString(list))

}
