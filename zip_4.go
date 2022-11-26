package lq

type zip4Iterator[T1, T2, T3, T4 any] struct {
	iterator1 Iterator[T1]
	iterator2 Iterator[T2]
	iterator3 Iterator[T3]
	iterator4 Iterator[T4]
}

func Zip4[T1, T2, T3, T4 any](
	iterator1 Iterator[T1],
	iterator2 Iterator[T2],
	iterator3 Iterator[T3],
	iterator4 Iterator[T4],
) Iterator[Tuple4[T1, T2, T3, T4]] {
	return zip4Iterator[T1, T2, T3, T4]{
		iterator1: iterator1,
		iterator2: iterator2,
		iterator3: iterator3,
		iterator4: iterator4,
	}
}

func (it zip4Iterator[T1, T2, T3, T4]) Count() int {
	count := Max(
		DefaultIfEmpty(
			Where(
				Values(
					tryEstimateCount(it.iterator1),
					tryEstimateCount(it.iterator2),
					tryEstimateCount(it.iterator3),
					tryEstimateCount(it.iterator4),
				),
				func(v int) bool { return v > 0 },
			),
		),
	)

	return count
}

func (it zip4Iterator[T1, T2, T3, T4]) Range(f func(v Tuple4[T1, T2, T3, T4]) bool) {
	ch1 := make(chan T1)
	ch2 := make(chan T2)
	ch3 := make(chan T3)
	ch4 := make(chan T4)
	done := make(chan struct{})

	go iterateIteratorToChannel(it.iterator1, ch1, done)
	go iterateIteratorToChannel(it.iterator2, ch2, done)
	go iterateIteratorToChannel(it.iterator3, ch3, done)
	go iterateIteratorToChannel(it.iterator4, ch4, done)

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

		v3, ok := <-ch3
		if !ok {
			done <- struct{}{}
			return
		}

		v4, ok := <-ch4
		if !ok {
			done <- struct{}{}
			return
		}

		if !f(NewTuple4(v1, v2, v3, v4)) {
			done <- struct{}{}
			return
		}
	}
}
