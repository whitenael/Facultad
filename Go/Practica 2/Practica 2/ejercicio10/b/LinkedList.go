package main

import "fmt"

type List[T any] interface {
}

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
}

func NewList[T any]() LinkedList[T] {
	return LinkedList[T]{}
}

func (self LinkedList[T]) IsEmpty() bool {
	return self.head == nil
}

func (self LinkedList[T]) Len() int {
	length := 0
	current := self.head
	for current != nil {
		length++
		current = current.next
	}

	return length
}

func (self LinkedList[T]) FrontElement() T {
	return self.head.data
}

func (self LinkedList[T]) LastElement() T {
	if self.IsEmpty() {
		var zero T
		return zero
	}

	current := self.head
	for current.next != nil {
		current = current.next
	}

	return current.data
}

func (self LinkedList[T]) Next() LinkedList[T] {
	return LinkedList[T]{self.head.next}
}

func (self LinkedList[T]) ToString() string {
	str := ""
	current := self.head
	for current != nil {
		str += fmt.Sprint("", current.data) + "-> "
		current = current.next
	}

	return str
}

func (self *LinkedList[T]) PushFront(elem T) {
	node := &Node[T]{data: elem, next: self.head}
	self.head = node
}

func (self *LinkedList[T]) PushBack(elem T) {
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

func (self *LinkedList[T]) Remove() T {
	if self.head == nil {
		var zero T
		return zero
	}

	data := self.head.data
	self.head = self.head.next
	return data
}

func (self LinkedList[T]) Iterate(f func(int) int) {
	current := self.head
	for current != nil {

	}
}
