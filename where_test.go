package lq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Where(t *testing.T) {
	t.Run(
		"should filter slice", func(t *testing.T) {
			values := Slice([]int{1, 2, 3}).
				Where(func(v int) bool { return v%2 == 1 }).
				ToSlice()

			require.Equal(t, []int{1, 3}, values)
		},
	)
}
