package lq

type distinctIterator[T comparable] struct {
	iterator Iterator[T]
}

func Distinct[T comparable](
	iterator Iterator[T],
) Iterator[T] {
	return distinctIterator[T]{
		iterator: iterator,
	}
}

func (it distinctIterator[T]) Range(f func(value T) bool) {
	visited := make(map[T]struct{})

	it.iterator.Range(
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
}
