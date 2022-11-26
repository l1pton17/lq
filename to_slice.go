package lq

func ToSlice[T any](iterator Iterator[T]) []T {
	values := make([]T, 0, tryEstimateCount(iterator))

	iterator.Range(
		func(value T) bool {
			values = append(values, value)
			return true
		},
	)

	return values
}
