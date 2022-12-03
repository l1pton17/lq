package lq

func Map[K comparable, V any](value map[K]V) Iterator[MapEntry[K, V]] {
	return Iterator[MapEntry[K, V]]{
		cheapCountFn: func() int {
			return len(value)
		},
		rangeFn: func(fn Iteratee[MapEntry[K, V]]) {
			for k, v := range value {
				if !fn(NewMapEntry(k, v)) {
					break
				}
			}
		},
	}
}
