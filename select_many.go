package lq

type selectManyIterator[TIn, TOut any] struct {
	iterator Iterator[TIn]
	selector func(value TIn) Iterator[TOut]
}

func SelectMany[TIn, TOut any](
	iterator Iterator[TIn],
	selector func(value TIn) Iterator[TOut],
) Iterator[TOut] {
	return selectManyIterator[TIn, TOut]{
		iterator: iterator,
		selector: selector,
	}
}

func (it selectManyIterator[TIn, TOut]) Range(f func(value TOut) bool) {
	stopped := false

	it.iterator.Range(
		func(v TIn) bool {
			it.selector(v).Range(
				func(ov TOut) bool {
					if !f(ov) {
						stopped = true
						return false
					}
					return true
				},
			)
			
			if stopped {
				return false
			}

			return true
		},
	)
}
