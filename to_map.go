package lq

func ToMap[TIn any, TKey comparable, TValue any](
	iterator Iterator[TIn],
	keySelector func(value TIn) TKey,
	valueSelector func(value TIn) TValue,
) map[TKey]TValue {
	values := make(map[TKey]TValue, tryEstimateCount(iterator))

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
