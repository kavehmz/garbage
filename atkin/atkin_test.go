package atkin

import "testing"

func BenchmarkPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Atkin(100000000)
	}
}
