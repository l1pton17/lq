package lq

func (it Iterator[T]) DefaultIfEmpty() Iterator[T] {
	return DefaultIfEmpty(it)
}

func (it Iterator[T]) DefaultIfEmptyV(defaultValue T) Iterator[T] {
	return DefaultIfEmptyV(it, defaultValue)
}

func DefaultIfEmpty[T any](
	iterator Iterator[T],
) Iterator[T] {
	var defaultValue T
	return DefaultIfEmptyV(iterator, defaultValue)
}

func DefaultIfEmptyV[T any](
	iterator Iterator[T],
	defaultValue T,
) Iterator[T] {
	return Iterator[T]{
		cheapCountFn: func() int {
			count := iterator.CheapCount()
			if count == 0 {
				return 1
			}

			return count
		},
		rangeFn: func(f Iteratee[T]) {
			hasValue := false

			iterator.Range(
				func(v T) bool {
					hasValue = true
					return f(v)
				},
			)

			if !hasValue {
				f(defaultValue)
			}
		},
	}
}
