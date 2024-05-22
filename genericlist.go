package main

import "errors"

type GenericList[T comparable] struct {
	data []T
}

func NewList[T comparable]() *GenericList[T] {
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

func (l *GenericList[T]) FirstIndex(value T) (int, bool) {
	index := -1
	found := false

	for i := 0; i < len(l.data); i++ {
		if value == l.data[i] {
			index = i
			found = true
			break
		}
	}

	return index, found
}

func (l *GenericList[T]) RemoveByValue(value T) (int, bool) {
	index, ok := l.FirstIndex(value)

	if !ok || index < 0 {
		return -1, false
	}

	l.data = append(l.data[:index], l.data[index+1:]...)
	return index, true
}
