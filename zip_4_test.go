package lq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Zip4(t *testing.T) {
	actual := ToSlice(
		Zip4(
			Range(1, 10),
			Range(5, 3),
			Range(5, 2),
			Range(7, 3),
		),
	)

	require.Equal(
		t,
		[]Tuple4[int, int, int, int]{
			T4(1, 5, 5, 7),
			T4(2, 6, 6, 8),
		},
		actual,
	)
}
