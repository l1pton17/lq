package lq

func Distinct[T comparable](
	iterator Iterator[T],
) Iterator[T] {
	return Iterator[T]{
		cheapCountFn: unknownCountFn,
		rangeFn: func(f Iteratee[T]) {
			visited := make(map[T]struct{})

			iterator.Range(
				func(value T) bool {
					if _, exists := visited[value]; !exists {
						if !f(value) {
							return false
						}

						visited[value] = struct{}{}
					}

					return true
				},
			)
		},
	}
}
