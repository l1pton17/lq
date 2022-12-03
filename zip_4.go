package lq

func Zip4[TA, TB, TC, TD any](
	iterator1 Iterator[TA],
	iterator2 Iterator[TB],
	iterator3 Iterator[TC],
	iterator4 Iterator[TD],
) Iterator[Tuple4[TA, TB, TC, TD]] {
	return Iterator[Tuple4[TA, TB, TC, TD]]{
		cheapCountFn: func() int {
			count := Max(
				Values(iterator1.CheapCount(), iterator2.CheapCount(), iterator3.CheapCount(), iterator4.CheapCount()).
					Where(func(v int) bool { return v > 0 }).
					DefaultIfEmpty(),
			)

			return count
		},
		rangeFn: func(f Iteratee[Tuple4[TA, TB, TC, TD]]) {
			ch1 := make(chan TA)
			ch2 := make(chan TB)
			ch3 := make(chan TC)
			ch4 := make(chan TD)
			done := make(chan struct{})

			go iterateIteratorToChannel(iterator1, ch1, done)
			go iterateIteratorToChannel(iterator2, ch2, done)
			go iterateIteratorToChannel(iterator3, ch3, done)
			go iterateIteratorToChannel(iterator4, ch4, done)

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

				if !f(T4(v1, v2, v3, v4)) {
					done <- struct{}{}
					return
				}
			}
		},
	}
}
