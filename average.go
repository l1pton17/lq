package lq

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

	if count == 0 {
		panic("no elements")
	}

	return float64(sum) / float64(count)
}
