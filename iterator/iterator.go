package iterator

import "rhin/lib"

type Iterator[T any] struct {
	items []T
}

func NewIterator[T any](items []T) *Iterator[T] {
	return &Iterator[T]{items}
}

func (iter Iterator[T]) Filter(condition func(item T) bool) *Iterator[T] {
	return NewIterator(lib.Filter(iter.items, condition))
}

func (iter Iterator[T]) Map(x interface{}, transform func(item T) any) *Iterator[any] {
	return NewIterator(lib.Map(iter.items, transform))
}
