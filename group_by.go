package lq

type GroupByEntry[TKey comparable, TValue any] struct {
	Key    TKey
	Values []TValue
}

type groupByIterator[TKey comparable, TValue any] struct {
	iterator    Iterator[TValue]
	keySelector func(value TValue) TKey
}

func GroupBy[TValue any, TKey comparable](
	iterator Iterator[TValue],
	keySelector func(value TValue) TKey,
) Iterator[GroupByEntry[TKey, TValue]] {
	return groupByIterator[TKey, TValue]{
		iterator:    iterator,
		keySelector: keySelector,
	}
}

func (g groupByIterator[TKey, TValue]) Range(
	f func(value GroupByEntry[TKey, TValue]) bool,
) {
	groups := make(map[TKey][]TValue)

	g.iterator.Range(
		func(value TValue) bool {
			key := g.keySelector(value)
			groups[key] = append(groups[key], value)
			return true
		},
	)

	for key, values := range groups {
		if !f(GroupByEntry[TKey, TValue]{Key: key, Values: values}) {
			return
		}
	}
}
