package lq

type counterIterator[T any] struct {
	iterator Iterator[T]
}

func Count[T any](iterator Iterator[T]) int {
	estimatedCount := tryEstimateCount(iterator)

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
