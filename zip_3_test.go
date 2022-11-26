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
			NewTuple3(1, 5, 5),
			NewTuple3(2, 6, 6),
		},
		actual,
	)
}
