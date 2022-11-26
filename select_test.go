package lq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Select(t *testing.T) {
	values := ToSlice(
		Select(
			Slice([]int{1, 2, 3}),
			func(v int) int {
				return v * 2
			},
		),
	)

	require.Equal(t, []int{2, 4, 6}, values)
}
