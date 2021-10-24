package demo

import "testing"

func BenchmarkAppendOne(b *testing.B) {
	num := 100000
	for i := 0; i < b.N; i++ {
		_ = appendOne(num)
	}
}

func BenchmarkAppendMany(b *testing.B) {
	num := 100000
	for i := 0; i < b.N; i++ {
		_ = appendMany(num)
	}
}
