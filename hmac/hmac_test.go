package hmac

import (
	"fmt"
	"testing"
)

var result string

func BenchmarkRun(b *testing.B) {
	result = Run()
	b.Log(result)
}

func BenchmarkRun100Times(b *testing.B) {
	for i := 1; i <= 100; i++ {
		Run()
	}
}

func BenchmarkAuth(b *testing.B) {
	b.Log(Auth(result))
}

func BenchmarkAuth100Times(b *testing.B) {
	for i := 1; i <= 100; i++ {
		Auth(result)
	}
}

func ExampleRun() {
	fmt.Println(Run())
}
