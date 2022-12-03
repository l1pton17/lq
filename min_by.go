package lq

func MinBy[TIn any, TOrderKey Ordered](
	iterator Iterator[TIn],
	selector func(v TIn) TOrderKey,
) TIn {
	var minValue TOrderKey
	var minElement TIn
	var minSet bool

	iterator.Range(
		func(value TIn) bool {
			if !minSet {
				minElement = value
				minValue = selector(value)
				minSet = true
			} else {
				orderKey := selector(value)

				if minValue > orderKey {
					minValue = orderKey
					minElement = value
				}
			}

			return true
		},
	)

	if !minSet {
		panic("no elements found")
	}

	return minElement
}
