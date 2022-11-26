package lq

type SliceEntry[T any] struct {
	Value T
	Index int
}

func NewSliceEntry[T any](index int, value T) SliceEntry[T] {
	return SliceEntry[T]{
		Value: value,
		Index: index,
	}
}
