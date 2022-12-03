package lq

type GroupByEntry[K comparable, V any] struct {
	Key    K
	Values []V
}

func GroupBy[K comparable, V any](
	iterator Iterator[V],
	keySelector func(v V) K,
) Iterator[GroupByEntry[K, V]] {
	return Iterator[GroupByEntry[K, V]]{
		cheapCountFn: unknownCountFn,
		rangeFn: func(f Iteratee[GroupByEntry[K, V]]) {
			groups := make(map[K][]V)

			iterator.Range(
				func(value V) bool {
					key := keySelector(value)
					groups[key] = append(groups[key], value)
					return true
				},
			)

			for key, values := range groups {
				if !f(GroupByEntry[K, V]{Key: key, Values: values}) {
					return
				}
			}
		},
	}
}
