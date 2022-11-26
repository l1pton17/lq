package lq

type reverseIterator[T any] struct {
	iterator Iterator[T]
}

func Reverse[T any](iterator Iterator[T]) Iterator[T] {
	return reverseIterator[T]{iterator: iterator}
}

func (it reverseIterator[T]) Count() int {
	return tryEstimateCount(it.iterator)
}

func (it reverseIterator[T]) Range(f func(value T) bool) {
	values := ToSlice(it.iterator)

	for i := len(values) - 1; i >= 0; i-- {
		if !f(values[i]) {
			return
		}
	}
}
