package lq

type whereIterator[T any] struct {
	iterator  Iterator[T]
	predicate func(value T) bool
}

func Where[T any](
	iterator Iterator[T],
	predicate func(v T) bool,
) Iterator[T] {
	return whereIterator[T]{
		iterator:  iterator,
		predicate: predicate,
	}
}

func (w whereIterator[T]) Range(f func(v T) bool) {
	w.iterator.Range(
		func(value T) bool {
			if w.predicate(value) {
				if !f(value) {
					return false
				}
			}

			return true
		},
	)
}
