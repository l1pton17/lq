package lq

const UnknownCount = 0

var unknownCountFn = func() int { return UnknownCount }
var zeroCountFn = func() int { return 0 }

type Iteratee[T any] func(v T) bool

type Iterator[T any] struct {
	rangeFn      func(f Iteratee[T])
	cheapCountFn func() int
}

func NewIterator[T any](rangeFn func(f Iteratee[T]), cheapCountFn func() int) Iterator[T] {
	return Iterator[T]{
		rangeFn:      rangeFn,
		cheapCountFn: cheapCountFn,
	}
}

func (it Iterator[T]) Range(f Iteratee[T]) {
	it.rangeFn(f)
}

func (it Iterator[T]) CheapCount() int {
	return it.cheapCountFn()
}
