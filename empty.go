package lq

type emptyIterator[T any] struct {
}

func Empty[T any]() Iterator[T] {
	return emptyIterator[T]{}
}

func (it emptyIterator[T]) Count() int {
	return 0
}

func (it emptyIterator[T]) Range(f func(v T) bool) {
}
