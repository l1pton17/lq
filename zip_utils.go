package lq

func iterateIteratorToChannel[T any](
	iterator Iterator[T],
	outCh chan<- T,
	done <-chan struct{},
) {
	iterator.Range(
		func(v T) bool {
			select {
			case outCh <- v:
			case <-done:
				return false
			}

			return true
		},
	)

	close(outCh)
}
