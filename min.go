package lq

func Min[T Ordered](iterator Iterator[T]) T {
	return MinBy(iterator, func(v T) T { return v })
}
