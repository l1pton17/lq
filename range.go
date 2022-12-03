package lq

func Range[T Math](min T, count int) Iterator[T] {
	return Iterator[T]{
		cheapCountFn: func() int {
			return count
		},
		rangeFn: func(f Iteratee[T]) {
			v := min

			for i := 0; i < count; i++ {
				if !f(v) {
					return
				}
				v++
			}
		},
	}
}
