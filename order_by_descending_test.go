package lq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_OrderByDescending_Compare(t *testing.T) {
	t.Parallel()

	t.Run(
		"should return 0 if values equal", func(t *testing.T) {
			subject := OrderByDescending(func(v int) int { return -1 * v })

			actual := subject.Compare(1, 1)

			require.Equal(t, 0, actual)
		},
	)

	t.Run(
		"should return -1 if value a greater than b", func(t *testing.T) {
			subject := OrderByDescending(func(v int) int { return v })

			actual := subject.Compare(2, 1)

			require.Equal(t, -1, actual)
		},
	)

	t.Run(
		"should return 1 if value a less than b", func(t *testing.T) {
			subject := OrderByDescending(func(v int) int { return v })

			actual := subject.Compare(0, 1)

			require.Equal(t, 1, actual)
		},
	)
}
