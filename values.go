package lq

func Values[T any](values ...T) Iterator[T] {
	return Iterator[T]{
		cheapCountFn: func() int {
			return len(values)
		},
		rangeFn: func(f Iteratee[T]) {
			for i := 0; i < len(values); i++ {
				if !f(values[i]) {
					return
				}
			}
		},
	}
}
