package lq

func (it Iterator[T]) Concat(
	iterator1 Iterator[T],
	iterators ...Iterator[T],
) Iterator[T] {
	return Concat(it, iterator1, iterators...)
}

func Concat[T any](
	iterator1 Iterator[T],
	iterator2 Iterator[T],
	iterators ...Iterator[T],
) Iterator[T] {
	return Iterator[T]{
		cheapCountFn: unknownCountFn,
		rangeFn: func(f Iteratee[T]) {
			stopped := false

			rit := func(value T) bool {
				if !f(value) {
					stopped = true
					return false
				}
				return true
			}

			iterator1.Range(rit)
			if stopped {
				return
			}

			iterator2.Range(rit)
			if stopped {
				return
			}

			for _, iterator := range iterators {
				iterator.Range(rit)
				if stopped {
					return
				}
			}
		},
	}
}
