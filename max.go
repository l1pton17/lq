package lq

func Max[T Ordered](iterator Iterator[T]) T {
	return MaxBy(iterator, func(v T) T { return v })
}
