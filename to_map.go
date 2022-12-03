package lq

func ToMap[TIn any, K comparable, V any](
	iterator Iterator[TIn],
	keySelector func(value TIn) K,
	valueSelector func(value TIn) V,
) map[K]V {
	values := make(map[K]V, iterator.CheapCount())

	iterator.Range(
		func(value TIn) bool {
			mKey := keySelector(value)
			mValue := valueSelector(value)
			values[mKey] = mValue

			return true
		},
	)

	return values
}
