package lq

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
