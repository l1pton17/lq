package lq

type zip3Iterator[T1, T2, T3 any] struct {
	iterator1 Iterator[T1]
	iterator2 Iterator[T2]
	iterator3 Iterator[T3]
}

func Zip3[T1, T2, T3 any](
	iterator1 Iterator[T1],
	iterator2 Iterator[T2],
	iterator3 Iterator[T3],
) Iterator[Tuple3[T1, T2, T3]] {
	return zip3Iterator[T1, T2, T3]{
		iterator1: iterator1,
		iterator2: iterator2,
		iterator3: iterator3,
	}
}

func (it zip3Iterator[T1, T2, T3]) Count() int {
	count := Max(
		DefaultIfEmpty(
			Where(
				Values(
					tryEstimateCount(it.iterator1),
					tryEstimateCount(it.iterator2),
					tryEstimateCount(it.iterator3),
				),
				func(v int) bool { return v > 0 },
			),
		),
	)

	return count
}

func (it zip3Iterator[T1, T2, T3]) Range(f func(v Tuple3[T1, T2, T3]) bool) {
	ch1 := make(chan T1)
	ch2 := make(chan T2)
	ch3 := make(chan T3)
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

	go func() {
		it.iterator3.Range(
			func(v T3) bool {
				select {
				case ch3 <- v:
				case <-done:
					return false
				}

				return true
			},
		)

		close(ch3)
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

		v3, ok := <-ch3
		if !ok {
			done <- struct{}{}
			return
		}

		if !f(NewTuple3(v1, v2, v3)) {
			done <- struct{}{}
			return
		}
	}
}
