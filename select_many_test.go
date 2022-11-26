package lq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SelectMany(t *testing.T) {
	actual := ToSlice(
		SelectMany(
			Range(1, 3),
			func(value int) Iterator[int] { return Repeat(value, 2) },
		),
	)

	require.Equal(t, []int{1, 1, 2, 2, 3, 3}, actual)
}

func Test_SelectMany_Where(t *testing.T) {
	actual := ToSlice(
		Where(
			SelectMany(
				Range(1, 4),
				func(value int) Iterator[int] { return Repeat(value, 2) },
			),
			func(v int) bool { return v%2 == 1 },
		),
	)

	require.Equal(t, []int{1, 1, 3, 3}, actual)
}
