package lq

type takeIterator[T any] struct {
	iterator Iterator[T]
	count    int
}

func Take[T any](
	iterator Iterator[T],
	count int,
) Iterator[T] {
	return takeIterator[T]{
		iterator: iterator,
		count:    count,
	}
}

func (it takeIterator[T]) Count() int {
	itCount := tryEstimateCount(it.iterator)

	if itCount < it.count {
		return itCount
	}
	return it.count
}

func (it takeIterator[T]) Range(f func(value T) bool) {
	count := 0

	it.iterator.Range(
		func(value T) bool {
			if count < it.count {
				if !f(value) {
					return false
				}

				count++

				return true
			}

			return false
		},
	)
}
