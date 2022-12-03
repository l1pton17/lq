package lq

func emptyRangeFn[T any](Iteratee[T]) {
}

func Empty[T any]() Iterator[T] {
	return Iterator[T]{
		cheapCountFn: zeroCountFn,
		rangeFn:      emptyRangeFn[T],
	}
}
