package lq

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Concat(t *testing.T) {
	t.Run(
		"should concat iterators", func(t *testing.T) {
			actual := Concat(
				Values(1, 2, 3),
				Values(4, 5, 6),
				Values(7, 8, 9),
			).ToSlice()

			require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, actual)
		},
	)

	t.Run(
		"should stop iterator", func(t *testing.T) {
			expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

			for i := 0; i < len(expected); i++ {
				t.Run(
					fmt.Sprintf("should stop iterator on %v", i), func(t *testing.T) {
						actual :=
							Concat(
								Values(1, 2, 3),
								Values(4, 5, 6),
								Values(7, 8, 9),
							).Take(i).ToSlice()

						require.Equal(t, expected[:i], actual)
					},
				)
			}
		},
	)
}
