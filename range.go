package lq

type rangeIterator[T Math] struct {
	min   T
	count int
}

func Range[T Math](min T, count int) Iterator[T] {
	return rangeIterator[T]{min: min, count: count}
}

func (it rangeIterator[T]) Count() int {
	return it.count
}

func (it rangeIterator[T]) Range(f func(value T) bool) {
	v := it.min

	for i := 0; i < it.count; i++ {
		if !f(v) {
			return
		}
		v++
	}
}
