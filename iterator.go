package lq

type Iterator[T any] interface {
	Range(func(v T) bool)
}

type countedIterator[T any] interface {
	Count() int
}

func tryEstimateCount[T any](iterator Iterator[T]) int {
	if ci, ok := iterator.(countedIterator[T]); ok {
		return ci.Count()
	}

	return 0
}
