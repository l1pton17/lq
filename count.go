package lq

func (it Iterator[T]) Count() int {
	return Count(it)
}

func Count[T any](iterator Iterator[T]) int {
	estimatedCount := iterator.CheapCount()
	if estimatedCount != 0 {
		return estimatedCount
	}

	count := 0

	iterator.Range(
		func(value T) bool {
			count++
			return true
		},
	)

	return count
}
