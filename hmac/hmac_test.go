package hmac

import (
	"testing"
)

func BenchmarkRun(b *testing.B) {
	b.Log(Run())
}
