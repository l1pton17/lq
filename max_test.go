package lq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Max(t *testing.T) {
	t.Run(
		"should return max element", func(t *testing.T) {
			actual := Max(Slice([]int{1, 15, 3, -5, 2}))
			require.Equal(t, 15, actual)
		},
	)
}
