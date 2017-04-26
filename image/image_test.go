package image

import (
	"testing"
)

func BenchmarkRun(b *testing.B) {
	_, err := Run()
	if err != nil {
		b.Error(err)
	}
	b.Logf("get")
}

func BenchmarkRunOverHttp(b *testing.B) {
	_ = RunOverHttp()
	b.Logf("end")
}
