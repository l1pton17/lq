package lq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Distinct(t *testing.T) {
	t.Run(
		"should return distinct values", func(t *testing.T) {
			actual := ToSlice(Distinct(Slice([]int{1, 2, 3, 1, 4, 5, 2, 3})))

			require.Equal(t, []int{1, 2, 3, 4, 5}, actual)
		},
	)
}
