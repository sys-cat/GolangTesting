package jwt

import "testing"

func BenchmarkRun(b *testing.B) {
	for i := 0; i < 10; i++ {
		b.Logf("%d times: %s", i, Run())
	}
}
