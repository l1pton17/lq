package lq

func (it Iterator[T]) Reverse() Iterator[T] {
	return Reverse(it)
}

func Reverse[T any](iterator Iterator[T]) Iterator[T] {
	return Iterator[T]{
		cheapCountFn: iterator.cheapCountFn,
		rangeFn: func(f Iteratee[T]) {
			values := ToSlice(iterator)

			for i := len(values) - 1; i >= 0; i-- {
				if !f(values[i]) {
					return
				}
			}
		},
	}
}
