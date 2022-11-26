package lq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Range(t *testing.T) {
	actual := ToSlice(Range(10, 5))

	require.Equal(t, []int{10, 11, 12, 13, 14}, actual)
}
