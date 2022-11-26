package lq

type concatIterator[T any] struct {
	iterator1 Iterator[T]
	iterator2 Iterator[T]
	iterators []Iterator[T]
}

func Concat[T any](
	iterator1 Iterator[T],
	iterator2 Iterator[T],
	iterators ...Iterator[T],
) Iterator[T] {
	return concatIterator[T]{
		iterator1: iterator1,
		iterator2: iterator2,
		iterators: iterators,
	}
}

func (s concatIterator[T]) Range(f func(value T) bool) {
	stopped := false

	rit := func(value T) bool {
		if !f(value) {
			stopped = true
			return false
		}
		return true
	}

	s.iterator1.Range(rit)
	if stopped {
		return
	}

	s.iterator2.Range(rit)
	if stopped {
		return
	}

	for _, iterator := range s.iterators {
		iterator.Range(rit)
		if stopped {
			return
		}
	}
}
