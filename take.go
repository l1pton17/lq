package lq

func (it Iterator[T]) Take(count int) Iterator[T] {
	return Take(it, count)
}

func Take[T any](
	iterator Iterator[T],
	count int,
) Iterator[T] {
	return Iterator[T]{
		cheapCountFn: func() int {
			itCount := iterator.CheapCount()
			if itCount < count {
				return itCount
			}

			return count
		},
		rangeFn: func(f Iteratee[T]) {
			curCount := 0

			iterator.Range(
				func(value T) bool {
					if curCount < count {
						curCount++

						return f(value)
					}

					return false
				},
			)
		},
	}
}
