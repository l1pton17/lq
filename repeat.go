package lq

type repeatIterator[T any] struct {
	value T
	count int
}

func Repeat[T any](value T, count int) Iterator[T] {
	return repeatIterator[T]{
		value: value,
		count: count,
	}
}

func (r repeatIterator[T]) Count() int {
	return r.count
}

func (r repeatIterator[T]) Range(f func(value T) bool) {
	for i := 0; i < r.count; i++ {
		if !f(r.value) {
			return
		}
	}
}
