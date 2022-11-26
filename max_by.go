package lq

func MaxBy[TIn any, TOut Ordered](
	iterator Iterator[TIn],
	selector func(v TIn) TOut,
) TOut {
	var max TOut
	var maxSet bool

	iterator.Range(
		func(value TIn) bool {
			if !maxSet {
				max = selector(value)
				maxSet = true
			} else {
				v := selector(value)

				if max < v {
					max = v
				}
			}

			return true
		},
	)

	if !maxSet {
		panic("no elements found")
	}

	return max
}
