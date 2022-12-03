package lq

type zip2Iterator[T1, T2 any] struct {
	iterator1 Iterator[T1]
	iterator2 Iterator[T2]
}

func Zip2[T1, T2 any](
	iterator1 Iterator[T1],
	iterator2 Iterator[T2],
) Iterator[Tuple2[T1, T2]] {
	return zip2Iterator[T1, T2]{
		iterator1: iterator1,
		iterator2: iterator2,
	}
}

func (it zip2Iterator[T1, T2]) Count() int {
	count := Max(
		DefaultIfEmpty(
			Where(
				Values(
					tryEstimateCount(it.iterator1),
					tryEstimateCount(it.iterator2),
				),
				func(v int) bool { return v > 0 },
			),
		),
	)

	return count
}

func (it zip2Iterator[TA, TB]) Range(f func(v Tuple2[TA, TB]) bool) {
	ch1 := make(chan TA)
	ch2 := make(chan TB)
	done := make(chan struct{})

	go iterateIteratorToChannel(it.iterator1, ch1, done)
	go iterateIteratorToChannel(it.iterator2, ch2, done)

	for {
		v1, ok := <-ch1
		if !ok {
			done <- struct{}{}
			return
		}

		v2, ok := <-ch2
		if !ok {
			done <- struct{}{}
			return
		}

		if !f(T2(v1, v2)) {
			done <- struct{}{}
			return
		}
	}
}
