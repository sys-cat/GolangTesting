package jwt

import "testing"

var token map[string]string

func BenchmarkRun(b *testing.B) {
	token = Run()
	b.Log(token)
}

func BenchmarkRun100Times(b *testing.B) {
	for i := 1; i <= 100; i++ {
		Run()
	}
}

func BenchmarkDeRun(b *testing.B) {
	b.Log(DeRun([]string{token["body"], token["verify"]}))
}

/*
func BenchmarkDecodeOrigin(b *testing.B) {
	b.Log(DecodeOrigin([]string{token["body"], token["verify"]}))
}
*/
