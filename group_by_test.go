package lq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GroupBy(t *testing.T) {
	type MultiValue struct {
		A int
		B int
	}

	actual := ToMap(
		GroupBy(
			Slice(
				[]MultiValue{
					{A: 1, B: 1},
					{A: 1, B: 2},
					{A: 2, B: 3},
					{A: 2, B: 1},
				},
			),
			func(value MultiValue) int { return value.A },
		),
		func(value GroupByEntry[int, MultiValue]) int { return value.Key },
		func(value GroupByEntry[int, MultiValue]) []int {
			return ToSlice(Select(Slice(value.Values), func(value MultiValue) int { return value.B }))
		},
	)

	require.Len(t, actual, 2)
	require.Equal(t, []int{1, 2}, actual[1])
	require.Equal(t, []int{3, 1}, actual[2])
}
