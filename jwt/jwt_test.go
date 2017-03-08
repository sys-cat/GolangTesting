package jwt

import "testing"

func BenchmarkRun(b *testing.B) {
	b.Log(Run())
}

func BenchmarkRun100Times(b *testing.B) {
	for i := 1; i <= 100; i++ {
		Run()
	}
}
