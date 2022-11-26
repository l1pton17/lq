package lq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Repeat(t *testing.T) {
	actual := ToSlice(Repeat(10, 4))

	require.Equal(t, []int{10, 10, 10, 10}, actual)
}
