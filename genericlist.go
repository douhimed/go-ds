package main

type GenericList[T comparable] struct {
	data []T
	size int
}

func New[T comparable]() *GenericList[T] {
	return &GenericList[T]{
		data: []T{},
	}
}

func (l *GenericList[T]) Add(value T) {
	l.data = append(l.data, value)
	l.size++
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
	return l.size
}
