package lq

import "sort"

type Orderer[TValue any, TBy Ordered] interface {
	Compare(a, b TValue) int
}

type orderIterator[TValue any, TBy Ordered] struct {
	iterator Iterator[TValue]
	order1   Orderer[TValue, TBy]
	orderers []Orderer[TValue, TBy]
}

func Order[TValue any, TBy Ordered](
	iterator Iterator[TValue],
	order1 Orderer[TValue, TBy],
	orderers ...Orderer[TValue, TBy],
) Iterator[TValue] {
	return orderIterator[TValue, TBy]{
		iterator: iterator,
		order1:   order1,
		orderers: orderers,
	}
}

func (it orderIterator[TValue, TBy]) Count() int {
	return tryEstimateCount(it.iterator)
}

func (it orderIterator[TValue, TBy]) Range(f func(value TValue) bool) {
	values := ToSlice(it.iterator)

	sort.Slice(
		values, func(i, j int) bool {
			cmp := it.order1.Compare(values[i], values[j])
			if cmp != 0 {
				return cmp < 0
			}

			for _, orderer := range it.orderers {
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
}
