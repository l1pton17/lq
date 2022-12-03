package lq

func Zip2[TA, TB any](
	iterator1 Iterator[TA],
	iterator2 Iterator[TB],
) Iterator[Tuple2[TA, TB]] {
	return Iterator[Tuple2[TA, TB]]{
		cheapCountFn: func() int {
			count := Max(
				Values(iterator1.CheapCount(), iterator2.CheapCount()).
					Where(func(v int) bool { return v > 0 }).
					DefaultIfEmpty(),
			)

			return count
		},
		rangeFn: func(f Iteratee[Tuple2[TA, TB]]) {
			ch1 := make(chan TA)
			ch2 := make(chan TB)
			done := make(chan struct{})

			go iterateIteratorToChannel(iterator1, ch1, done)
			go iterateIteratorToChannel(iterator2, ch2, done)

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
		},
	}
}
