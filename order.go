package lq

import "sort"

type Orderer[TValue any, TBy Ordered] interface {
	Compare(a, b TValue) int
}

func Order[TValue any, TBy Ordered](
	iterator Iterator[TValue],
	order1 Orderer[TValue, TBy],
	orderers ...Orderer[TValue, TBy],
) Iterator[TValue] {
	return Iterator[TValue]{
		cheapCountFn: iterator.cheapCountFn,
		rangeFn: func(f Iteratee[TValue]) {
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
