package lq

func Select[TIn, TOut any](
	iterator Iterator[TIn],
	selector func(v TIn) TOut,
) Iterator[TOut] {
	return Iterator[TOut]{
		cheapCountFn: iterator.cheapCountFn,
		rangeFn: func(f Iteratee[TOut]) {
			iterator.Range(
				func(value TIn) bool {
					return f(selector(value))
				},
			)
		},
	}
}
