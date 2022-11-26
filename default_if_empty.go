package lq

type defaultIfEmptyIterator[T any] struct {
	iterator     Iterator[T]
	defaultValue T
}

func DefaultIfEmpty[T any](
	iterator Iterator[T],
) Iterator[T] {
	return defaultIfEmptyIterator[T]{
		iterator: iterator,
	}
}

func DefaultIfEmptyV[T any](
	iterator Iterator[T],
	defaultValue T,
) Iterator[T] {
	return defaultIfEmptyIterator[T]{
		iterator:     iterator,
		defaultValue: defaultValue,
	}
}

func (it defaultIfEmptyIterator[T]) Count() int {
	count := tryEstimateCount(it.iterator)

	if count == 0 {
		return 1
	}

	return count
}

func (it defaultIfEmptyIterator[T]) Range(f func(v T) bool) {
	hasValue := false

	it.iterator.Range(
		func(v T) bool {
			hasValue = true
			return f(v)
		},
	)

	if !hasValue {
		f(it.defaultValue)
	}
}
