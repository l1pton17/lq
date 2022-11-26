package lq

func MinBy[TIn any, TOut Ordered](
	iterator Iterator[TIn],
	selector func(v TIn) TOut,
) TOut {
	var min TOut
	var minSet bool

	iterator.Range(
		func(value TIn) bool {
			if !minSet {
				min = selector(value)
				minSet = true
			} else {
				v := selector(value)

				if min > v {
					min = v
				}
			}

			return true
		},
	)

	if !minSet {
		panic("no elements found")
	}

	return min
}
