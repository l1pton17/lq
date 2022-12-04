package lq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Order(t *testing.T) {
	type MultiOrder struct {
		A, B, C int
	}

	t.Run(
		"should order", func(t *testing.T) {
			actual := Slice([]int{3, 9, 2}).
				Order(OrderBy(func(value int) int { return value })).
				ToSlice()

			require.Equal(t, []int{2, 3, 9}, actual)
		},
	)

	t.Run(
		"should order by multiple times", func(t *testing.T) {
			values := []MultiOrder{
				{A: 1, B: 3, C: 3},
				{A: 1, B: 2, C: 3},
				{A: 1, B: 3, C: 4},
				{A: 1, B: 3, C: 3},
			}

			actual := Slice(values).
				Order(
					OrderBy(func(value MultiOrder) int { return value.A }),
					OrderBy(func(value MultiOrder) int { return value.B }),
					OrderBy(func(value MultiOrder) int { return value.C }),
				).ToSlice()

			require.Equal(
				t,
				[]MultiOrder{
					{A: 1, B: 2, C: 3},
					{A: 1, B: 3, C: 3},
					{A: 1, B: 3, C: 3},
					{A: 1, B: 3, C: 4},
				},
				actual,
			)
		},
	)

	t.Run(
		"should order by descending multiple times", func(t *testing.T) {
			values := []MultiOrder{
				{A: 1, B: 3, C: 3},
				{A: 1, B: 2, C: 3},
				{A: 1, B: 3, C: 4},
				{A: 1, B: 3, C: 3},
			}

			actual := Slice(values).
				Order(
					OrderByDesc(func(value MultiOrder) int { return value.A }),
					OrderByDesc(func(value MultiOrder) int { return value.B }),
					OrderByDesc(func(value MultiOrder) int { return value.C }),
				).ToSlice()

			require.Equal(
				t,
				[]MultiOrder{
					{A: 1, B: 3, C: 4},
					{A: 1, B: 3, C: 3},
					{A: 1, B: 3, C: 3},
					{A: 1, B: 2, C: 3},
				},
				actual,
			)
		},
	)
}
