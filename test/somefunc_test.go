package test

import (
	"fmt"
	"github.com/alexeyzcom/go-module/internal/service"
	"testing"
)

var table = []struct {
	input int
}{
	{input: 100},
	{input: 1000},
	{input: 74382},
	{input: 382399},
}

func BenchmarkPrimeNumbers(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				service.PrimeNumbers(v.input)
			}
		})
	}
}
