package lq

import "sort"

type Orderer[TValue any] interface {
	Compare(a, b TValue) int
}

func (it Iterator[T]) Order(
	order1 Orderer[T],
	orderers ...Orderer[T],
) Iterator[T] {
	return Order(it, order1, orderers...)
}

func Order[T any](
	iterator Iterator[T],
	order1 Orderer[T],
	orderers ...Orderer[T],
) Iterator[T] {
	return Iterator[T]{
		cheapCountFn: iterator.cheapCountFn,
		rangeFn: func(f Iteratee[T]) {
			values := ToSlice(iterator)

			sort.Slice(
				values, func(i, j int) bool {
					cmp := order1.Compare(values[i], values[j])
					if cmp != 0 {
						return cmp < 0
					}

					for _, orderer := range orderers {
						cmp = orderer.Compare(values[i], values[j])
						if cmp == 0 {
							continue
						} else {
							return cmp < 0
						}
					}

					return false
				},
			)

			for _, v := range values {
				if !f(v) {
					return
				}
			}
		},
	}
}
