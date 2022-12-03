package lq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Range(t *testing.T) {
	t.Run(
		"should return count value from min inclusive", func(t *testing.T) {
			actual := ToSlice(Range(10, 5))

			require.Equal(t, []int{10, 11, 12, 13, 14}, actual)
		},
	)

	t.Run(
		"should stop iterate if false returned", func(t *testing.T) {
			actual := ToSlice(Where(Range(10, 10), func(v int) bool { return v <= 11 }))
		},
	)
}
