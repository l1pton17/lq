package lq

func MaxBy[TIn any, TOrderBy Ordered](
	iterator Iterator[TIn],
	selector func(v TIn) TOrderBy,
) TIn {
	var maxValue TOrderBy
	var maxElement TIn
	var maxSet bool

	iterator.Range(
		func(value TIn) bool {
			if !maxSet {
				maxElement = value
				maxValue = selector(value)
				maxSet = true
			} else {
				orderer := selector(value)

				if maxValue < orderer {
					maxValue = orderer
					maxElement = value
				}
			}

			return true
		},
	)

	if !maxSet {
		panic("no elements found")
	}

	return maxElement
}
