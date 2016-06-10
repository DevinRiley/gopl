package popcount_test

import (
  "gopl/ch2/popcount"
  "gopl/ch2/2_3/popcountloop"
  "testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcountloop.PopCount(0x1234567890ABCDEF)
	}
}
