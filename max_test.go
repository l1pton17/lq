package lq

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Max(t *testing.T) {
	t.Parallel()

	t.Run(
		"should return max element", func(t *testing.T) {
			actual := Max(Slice([]int{1, 15, 3, -5, 2}))
			require.Equal(t, 15, actual)
		},
	)

	t.Run(
		"should panic if iterator is empty", func(t *testing.T) {
			assert.Panics(
				t, func() {
					Max(Slice([]int{}))
				},
			)
		},
	)
}
