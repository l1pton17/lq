package lq

func Slice[T any](value []T) Iterator[T] {
	return Iterator[T]{
		cheapCountFn: func() int {
			return len(value)
		},
		rangeFn: func(f Iteratee[T]) {
			for _, v := range value {
				if !f(v) {
					break
				}
			}
		},
	}
}
