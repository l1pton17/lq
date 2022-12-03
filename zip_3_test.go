package lq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Zip3(t *testing.T) {
	actual := ToSlice(
		Zip3(
			Range(1, 10),
			Range(5, 3),
			Range(5, 2),
		),
	)

	require.Equal(
		t,
		[]Tuple3[int, int, int]{
			T3(1, 5, 5),
			T3(2, 6, 6),
		},
		actual,
	)
}
