package lq

func SelectMany[TIn, TOut any](
	iterator Iterator[TIn],
	selector func(v TIn) Iterator[TOut],
) Iterator[TOut] {
	return Iterator[TOut]{
		cheapCountFn: func() int {
			return iterator.CheapCount()
		},
		rangeFn: func(f Iteratee[TOut]) {
			stopped := false

			iterator.Range(
				func(v TIn) bool {
					selector(v).Range(
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
		},
	}
}
