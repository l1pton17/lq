package lq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_DefaultIfEmpty(t *testing.T) {
	t.Run(
		"shouldn't return value when iterator isn't empty", func(t *testing.T) {
			actual := ToSlice(DefaultIfEmpty(Values(1, 2, 3)))

			require.Equal(t, []int{1, 2, 3}, actual)
		},
	)

	t.Run(
		"should return default value when iterator is empty", func(t *testing.T) {
			actual := ToSlice(DefaultIfEmpty(Empty[int]()))

			require.Equal(t, []int{0}, actual)
		},
	)

	t.Run(
		"should return passed default value when iterator is empty", func(t *testing.T) {
			actual := ToSlice(DefaultIfEmptyV(Empty[int](), -1))

			require.Equal(t, []int{-1}, actual)
		},
	)
}
