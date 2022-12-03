package lq

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/samber/lo"
)

func sliceGenerator(size uint) []int64 {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	result := make([]int64, size)

	for i := uint(0); i < size; i++ {
		result[i] = r.Int63()
	}

	return result
}

func BenchmarkSlice(b *testing.B) {
	arr := sliceGenerator(1_000_000)

	b.Run(
		"lq.Map", func(b *testing.B) {
			b.ReportAllocs()
			for n := 0; n < b.N; n++ {
				_ = Select(Slice(arr), func(v int64) string { return strconv.FormatInt(v, 10) }).
					Where(func(v string) bool { return len(v) < 1000 }).
					Where(func(v string) bool { return len(v) > 1 }).
					ToSlice()
			}
		},
	)

	b.Run(
		"lo.Map", func(b *testing.B) {
			b.ReportAllocs()
			for n := 0; n < b.N; n++ {
				_ = lo.Filter(
					lo.Filter(
						lo.Map(arr, func(x int64, i int) string { return strconv.FormatInt(x, 10) }),
						func(v string, i int) bool { return len(v) < 1000 },
					),
					func(v string, i int) bool { return len(v) > 1 },
				)
			}
		},
	)

	b.Run(
		"for", func(b *testing.B) {
			b.ReportAllocs()
			for n := 0; n < b.N; n++ {
				var results []string

				for _, item := range arr {
					result := strconv.FormatInt(item, 10)

					if len(result) < 1000 {
						results = append(results, result)
					}
				}
			}
		},
	)
}
