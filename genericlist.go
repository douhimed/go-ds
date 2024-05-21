package main

import "errors"

type GenericList[T comparable] struct {
	data []T
}

func New[T comparable]() *GenericList[T] {
	return &GenericList[T]{
		data: []T{},
	}
}

func (l *GenericList[T]) Add(value T) {
	l.data = append(l.data, value)
}

func (l *GenericList[T]) Get(index int) T {
	if index > len(l.data)-1 {
		panic("outside the range")
	}

	var value T
	for i := 0; i < len(l.data); i++ {
		if i == index {
			value = l.data[i]
			break
		}
	}
	return value
}

func (l *GenericList[T]) Length() int {
	return len(l.data)
}

func (l *GenericList[T]) Remove(i int) error {
	if i < 0 || i >= len(l.data) {
		return errors.New("out of range")
	}

	l.data = append(l.data[:i], l.data[i+1:]...)
	return nil
}
