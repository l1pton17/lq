package lq

func Repeat[T any](value T, count int) Iterator[T] {
	return Iterator[T]{
		cheapCountFn: func() int {
			return count
		},
		rangeFn: func(f Iteratee[T]) {
			for i := 0; i < count; i++ {
				if !f(value) {
					return
				}
			}
		},
	}
}
