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

func (it zip2Iterator[T1, T2]) Range(f func(v Tuple2[T1, T2]) bool) {
	ch1 := make(chan T1)
	ch2 := make(chan T2)
	done := make(chan struct{})

	go func() {
		it.iterator1.Range(
			func(v T1) bool {
				select {
				case ch1 <- v:
				case <-done:
					return false
				}

				return true
			},
		)

		close(ch1)
	}()

	go func() {
		it.iterator2.Range(
			func(v T2) bool {
				select {
				case ch2 <- v:
				case <-done:
					return false
				}

				return true
			},
		)

		close(ch2)
	}()

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

		if !f(NewTuple2(v1, v2)) {
			done <- struct{}{}
			return
		}
	}
}
