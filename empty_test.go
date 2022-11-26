package lq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Empty(t *testing.T) {
	t.Run(
		"should return zero count", func(t *testing.T) {
			actual := tryEstimateCount(Empty[int]())

			require.Equal(t, 0, actual)
		},
	)

	t.Run(
		"shouldn't enumerate function", func(t *testing.T) {
			Empty[int]().Range(
				func(v int) bool {
					require.FailNow(t, "enumeration happens")
					return true
				},
			)
		},
	)
}
