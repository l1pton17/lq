package lq

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_MinBy(t *testing.T) {
	t.Parallel()

	t.Run(
		"should return min by selector", func(t *testing.T) {
			actual := MinBy(
				Map(
					map[int]int{
						10: 4,
						3:  12,
					},
				),
				func(v MapEntry[int, int]) int { return v.Value },
			)

			require.Equal(t, 10, actual.Key)
			require.Equal(t, 4, actual.Value)
		},
	)

	t.Run(
		"should panic if iterator is empty", func(t *testing.T) {
			assert.Panics(
				t, func() {
					MinBy(
						Map(map[int]int{}),
						func(v MapEntry[int, int]) int { return v.Value },
					)
				},
			)
		},
	)
}
