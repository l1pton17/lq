package lq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_MapIterator_Count(t *testing.T) {
	t.Parallel()

	t.Run(
		"should return count of map", func(t *testing.T) {
			subject := Map(
				map[int]int{
					1: 2,
					3: 4,
				},
			)

			actual := subject.CheapCount()

			require.Equal(t, 2, actual)
		},
	)
}

func Test_MapIterator_Range(t *testing.T) {
	t.Parallel()

	t.Run(
		"should stop iterator if false returned", func(t *testing.T) {
			subject := map[int]int{
				1: 2,
				3: 4,
			}

			Map(subject).Range(
				func(v MapEntry[int, int]) bool {
					delete(subject, v.Key)
					return false
				},
			)

			require.Len(t, subject, 1)
		},
	)

	t.Run(
		"should iterate over map", func(t *testing.T) {
			expected := map[int]int{
				1: 2,
				3: 4,
			}

			Map(expected).Range(
				func(v MapEntry[int, int]) bool {
					require.Equal(t, expected[v.Key], v.Value)
					delete(expected, v.Key)

					return true
				},
			)

			require.Empty(t, expected)
		},
	)

	t.Run(
		"should iterate over custom map type", func(t *testing.T) {
			type CustomMap map[int]int

			expected := CustomMap{
				1: 2,
				3: 4,
			}

			Map(expected).Range(
				func(v MapEntry[int, int]) bool {
					require.Equal(t, expected[v.Key], v.Value)
					delete(expected, v.Key)

					return true
				},
			)

			require.Empty(t, expected)
		},
	)
}
