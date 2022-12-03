package lq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SelectMany(t *testing.T) {
	actual := SelectMany(
		Range(1, 3),
		func(v int) Iterator[int] { return Repeat(v, 2) },
	).ToSlice()

	require.Equal(t, []int{1, 1, 2, 2, 3, 3}, actual)
}

func Test_SelectMany_Where(t *testing.T) {
	actual := SelectMany(
		Range(1, 4),
		func(v int) Iterator[int] { return Repeat(v, 2) },
	).Where(func(v int) bool { return v%2 == 1 }).ToSlice()

	require.Equal(t, []int{1, 1, 3, 3}, actual)
}
