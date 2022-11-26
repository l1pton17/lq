package lq

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Average(t *testing.T) {
	t.Run(
		"should panic if iterator empty", func(t *testing.T) {
			assert.Panics(
				t,
				func() {
					_ = Average(Values[int]())
				},
			)
		},
	)

	t.Run(
		"should compute average", func(t *testing.T) {
			actual := Average(Values(1, 2, 3, 4))

			require.Equal(t, 2.5, actual)
		},
	)
}
