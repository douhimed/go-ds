package main

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
	size int
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (ll *LinkedList[T]) Length() int {
	return ll.size
}

func (ll *LinkedList[T]) Add(v T) {

	nn := &Node[T]{value: v}

	if ll.IsEmpty() {
		ll.head = nn
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = nn
	}

	ll.size++
}

func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.size == 0
}

func (ll *LinkedList[T]) IndexOf(v T) int {
	current := ll.head
	index := 0
	for current != nil {
		if current.value == v {
			return index
		}
		index++
		current = current.next
	}

	return -1
}
