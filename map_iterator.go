package lq

type MapIterator[TKey comparable, TValue any] struct {
	value map[TKey]TValue
}

func Map[TKey comparable, TValue any](value map[TKey]TValue) Iterator[MapEntry[TKey, TValue]] {
	return MapIterator[TKey, TValue]{value: value}
}

func (it MapIterator[TKey, TValue]) Count() int {
	return len(it.value)
}

func (it MapIterator[TKey, TValue]) Range(iterator func(entry MapEntry[TKey, TValue]) bool) {
	for k, v := range it.value {
		if !iterator(NewMapEntry(k, v)) {
			break
		}
	}
}
