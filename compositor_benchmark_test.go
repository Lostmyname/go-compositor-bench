package compositor

import (
	"testing"
)

const (
	PAGE_COUNT = 35
)

var assets = []string{
		"assets/001_small.png",
		"assets/002_small.png",
		"assets/003_small.png",
		"assets/004_small.png",
	}

func BenchmarkGeneratePage(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GeneratePage(assets, 1)
	}
}

func BenchmarkGeneratePagesSync(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GeneratePagesSync(assets, PAGE_COUNT)
	}
}

func BenchmarkGeneratePagesAsync(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GeneratePagesAsync(assets, PAGE_COUNT)
	}
}
