package lq

func (it Iterator[T]) ToSlice() []T {
	return ToSlice(it)
}

func ToSlice[T any](iterator Iterator[T]) []T {
	values := make([]T, 0, iterator.CheapCount())

	iterator.Range(
		func(value T) bool {
			values = append(values, value)
			return true
		},
	)

	return values
}
