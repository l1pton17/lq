package lq

type MapIterator[K comparable, V any] struct {
	value map[K]V
}

func Map[K comparable, V any](value map[K]V) Iterator[MapEntry[K, V]] {
	return MapIterator[K, V]{value: value}
}

func (it MapIterator[K, V]) Count() int {
	return len(it.value)
}

func (it MapIterator[K, V]) Range(iterator func(entry MapEntry[K, V]) bool) {
	for k, v := range it.value {
		if !iterator(NewMapEntry(k, v)) {
			break
		}
	}
}
