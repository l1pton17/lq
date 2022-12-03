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
			actual := ToSlice(
				Order(
					Slice([]int{3, 9, 2}),
					OrderBy(
						func(value int) int {
							return value
						},
					),
				),
			)

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

			actual := ToSlice(
				Order(
					Slice(values),
					OrderBy(func(value MultiOrder) int { return value.A }),
					OrderBy(func(value MultiOrder) int { return value.B }),
					OrderBy(func(value MultiOrder) int { return value.C }),
				),
			)

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

			actual := ToSlice(
				Order(
					Slice(values),
					OrderByDescending(func(value MultiOrder) int { return value.A }),
					OrderByDescending(func(value MultiOrder) int { return value.B }),
					OrderByDescending(func(value MultiOrder) int { return value.C }),
				),
			)

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
