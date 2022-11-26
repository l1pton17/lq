package lq

type valuesIterator[T any] struct {
	values []T
}

func Values[T any](values ...T) Iterator[T] {
	return valuesIterator[T]{
		values: values,
	}
}

func (it valuesIterator[T]) Count() int {
	return len(it.values)
}

func (it valuesIterator[T]) Range(f func(v T) bool) {
	for i := 0; i < len(it.values); i++ {
		if !f(it.values[i]) {
			return
		}
	}
}
