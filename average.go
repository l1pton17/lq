package lq

type averageIterator[T Math] struct {
	iterator Iterator[T]
}

func Average[T Math](iterator Iterator[T]) float64 {
	count := 0
	var sum T

	iterator.Range(
		func(value T) bool {
			sum += value
			count++

			return true
		},
	)

	return float64(sum) / float64(count)
}
