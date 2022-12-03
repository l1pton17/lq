package lq

func (it Iterator[T]) Skip(count int) Iterator[T] {
	return Skip(it, count)
}

func Skip[T any](
	iterator Iterator[T],
	count int,
) Iterator[T] {
	return Iterator[T]{
		cheapCountFn: func() int {
			itCount := iterator.CheapCount()
			if itCount > count {
				return 0
			}
			return itCount - count
		},
		rangeFn: func(f Iteratee[T]) {
			curCount := 0

			iterator.Range(
				func(value T) bool {
					if curCount >= count {
						return f(value)
					}

					curCount++

					return true
				},
			)
		},
	}
}
