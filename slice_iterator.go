package lq

type sliceIterator[T any] struct {
	value []T
}

func Slice[T any](value []T) Iterator[T] {
	return sliceIterator[T]{
		value: value,
	}
}

func (it sliceIterator[T]) Count() int {
	return len(it.value)
}

func (it sliceIterator[T]) Range(iterator func(value T) bool) {
	for _, v := range it.value {
		if !iterator(v) {
			break
		}
	}
}
