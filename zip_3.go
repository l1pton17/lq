package lq

func Zip3[TA, TB, TC any](
	iterator1 Iterator[TA],
	iterator2 Iterator[TB],
	iterator3 Iterator[TC],
) Iterator[Tuple3[TA, TB, TC]] {
	return Iterator[Tuple3[TA, TB, TC]]{
		cheapCountFn: func() int {
			count := Max(
				Values(iterator1.CheapCount(), iterator2.CheapCount(), iterator3.CheapCount()).
					Where(func(v int) bool { return v > 0 }).
					DefaultIfEmpty(),
			)

			return count
		},
		rangeFn: func(f Iteratee[Tuple3[TA, TB, TC]]) {
			ch1 := make(chan TA)
			ch2 := make(chan TB)
			ch3 := make(chan TC)
			done := make(chan struct{})

			go iterateIteratorToChannel(iterator1, ch1, done)
			go iterateIteratorToChannel(iterator2, ch2, done)
			go iterateIteratorToChannel(iterator3, ch3, done)

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

				if !f(T3(v1, v2, v3)) {
					done <- struct{}{}
					return
				}
			}
		},
	}
}
