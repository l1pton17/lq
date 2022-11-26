package lq

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func Test_Count(t *testing.T) {
	actual := Count(Repeat(10, 1_000_000))

	require.Equal(t, 1_000_000, actual)
}

func BenchmarkCount(b *testing.B) {
	b.Run(
		"lq.Count", func(b *testing.B) {
			b.ReportAllocs()
			for n := 0; n < b.N; n++ {
				_ = Count(
					Where(
						Range(1, 1_000_000),
						func(v int) bool { return v%2 == 0 },
					),
				)
			}
		},
	)

	b.Run(
		"lo.Count", func(b *testing.B) {
			b.ReportAllocs()
			for n := 0; n < b.N; n++ {
				_ = len(
					lo.Filter(
						lo.RangeFrom(1, 1_000_000),
						func(v int, i int) bool { return v%2 == 0 },
					),
				)
			}
		},
	)
}
