package lq

type skipIterator[T any] struct {
	iterator Iterator[T]
	count    int
}

func Skip[T any](
	iterator Iterator[T],
	count int,
) Iterator[T] {
	return skipIterator[T]{
		iterator: iterator,
		count:    count,
	}
}

func (it skipIterator[T]) Count() int {
	itCount := tryEstimateCount(it.iterator)
	if itCount > it.count {
		return 0
	}
	return itCount - it.count
}

func (it skipIterator[T]) Range(f func(value T) bool) {
	count := 0

	it.iterator.Range(
		func(value T) bool {
			if count > it.count {
				if !f(value) {
					return false
				}

				return true
			}

			count++

			return false
		},
	)
}
