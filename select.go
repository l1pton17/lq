package lq

type selectIterator[TIn, TOut any] struct {
	iterator Iterator[TIn]
	selector func(value TIn) TOut
}

func Select[TIn, TOut any](
	iterator Iterator[TIn],
	selector func(value TIn) TOut,
) Iterator[TOut] {
	return selectIterator[TIn, TOut]{
		iterator: iterator,
		selector: selector,
	}
}

func (s selectIterator[TIn, TOut]) Count() int {
	return tryEstimateCount(s.iterator)
}

func (s selectIterator[TIn, TOut]) Range(f func(value TOut) bool) {
	s.iterator.Range(
		func(value TIn) bool {
			outValue := s.selector(value)

			if !f(outValue) {
				return false
			}

			return true
		},
	)
}
