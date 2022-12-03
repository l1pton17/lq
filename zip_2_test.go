package lq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Zip2(t *testing.T) {
	actual := ToSlice(
		Zip2(
			Range(1, 10),
			Range(5, 3),
		),
	)

	require.Equal(
		t,
		[]Tuple2[int, int]{
			T2(1, 5),
			T2(2, 6),
			T2(3, 7),
		},
		actual,
	)
}
