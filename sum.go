package lq

type sumIterator[T Math] struct {
	iterator Iterator[T]
}

func Sum[T Math](iterator Iterator[T]) T {
	var sum T

	iterator.Range(
		func(value T) bool {
			sum += value
			return true
		},
	)

	return sum
}
