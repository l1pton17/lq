package lq

func (it Iterator[T]) Where(predicate func(v T) bool) Iterator[T] {
	return Where(it, predicate)
}

func Where[T any](
	iterator Iterator[T],
	predicate func(v T) bool,
) Iterator[T] {
	return Iterator[T]{
		cheapCountFn: unknownCountFn,
		rangeFn: func(f Iteratee[T]) {
			iterator.Range(
				func(v T) bool {
					if predicate(v) {
						return f(v)
					}
					return true
				},
			)
		},
	}
}
